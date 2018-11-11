package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/ordertype"
	storetypes "github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/store"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/symbol"
)

func New(opts *Options) (*Service, error) {
	db, err := sql.Open("postgres", opts.PostgresURL)
	if err != nil {
		log.Printf("err opening postgres; err: %v", err)
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("err creating driver\n%v", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///./migrations/",
		"postgres", driver,
	)
	if err != nil {
		log.Printf("err creating migration\n%v", err)
		return nil, err
	}
	m.Steps(1)

	return &Service{
		opts: opts,
		db:   db,
	}, db.Ping()
}

func (s *Service) GetBalance(get *storetypes.GetBalance) (*storetypes.GetBalanceReturn, error) {
	// note: balance has non-null constraint
	var balance string
	if err := s.db.QueryRow(`
		SELECT
			balances.balance::text

		FROM
			balances
			LEFT JOIN accounts ON balances.account_id = accounts.id

		WHERE
			UPPER(accounts.id) = UPPER($1) AND
			UPPER(balances.currency) = UPPER($2)

		LIMIT
			1;
	`, get.Account, get.Currency).Scan(&balance); err != nil {
		if err == sql.ErrNoRows {
			return &storetypes.GetBalanceReturn{
				Account:  get.Account,
				Currency: get.Currency,
				Balance:  big.NewInt(0),
			}, nil
		}

		log.Printf("err getting balance\n%v", err)
		return nil, err
	}

	i := new(big.Int)
	_, ok := i.SetString(balance, 10)
	if !ok {
		log.Println("err setting big int")
		return nil, errors.New("could not set big int")
	}

	return &storetypes.GetBalanceReturn{
		Account:  get.Account,
		Currency: get.Currency,
		Balance:  i,
	}, nil
}

func (s *Service) GetTradableBalance(get *storetypes.GetBalance) (*storetypes.GetBalanceReturn, error) {
	fullBalance, err := s.GetBalance(get)
	if err != nil {
		log.Printf("err getting balance\n%v", err)
		return nil, err
	}

	// note: balance has non-null constraint
	rows, err := s.db.Query(`
		SELECT
			UPPER(orderbook.symbol),
			UPPER(oderbook.type),
			orderbook.rate,
			orderbook.quantity::text

		FROM
			orderbook
			LEFT JOIN accounts ON orderbook.account_id = accounts.id

		WHERE
			UPPER(accounts.id) = UPPER($1)
	`, get.Account)
	if err != nil {
		if err == sql.ErrNoRows {
			return fullBalance, nil
		}

		log.Printf("err getting rows\n%v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			symbol, t, quantity string
			rate                float64
		)

		if err = rows.Scan(&symbol, &t, &rate, &quantity); err != nil {
			log.Printf("err scanning rows\n%v", err)
			return nil, err
		}

		coins := strings.Split(symbol, "_")
		if len(coins) != 2 {
			log.Printf("expected len(coins) == 2; received %d", len(coins))
			return nil, errors.New("coins len != 2")
		}

		if strings.ToUpper(coins[0]) == strings.ToUpper(get.Currency.String()) {
			// currency is base
			// only want asks
			if strings.ToUpper(t) == "ASK" {
				r := big.NewFloat(rate)

				q := new(big.Float)
				_, ok := q.SetString(quantity)
				if !ok {
					log.Printf("err setting big float\n%v", err)
					return nil, fmt.Errorf("could not set big float\n%v", err)
				}

				_ = q.Mul(q, r)

				b := new(big.Int)
				_, _ = q.Int(b)

				_ = fullBalance.Balance.Sub(fullBalance.Balance, b)
			}

		} else {
			// currency is quote
			// only want bids
			if strings.ToUpper(t) == "BID" {
				q := new(big.Int)
				_, ok := q.SetString(quantity, 10)
				if !ok {
					log.Printf("err setting big int\n%v", err)
					return nil, fmt.Errorf("could not set big int\n%v", err)
				}

				_ = fullBalance.Balance.Sub(fullBalance.Balance, q)
			}

		}
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows err\n%v", err)
		return nil, err
	}

	return fullBalance, nil
}

func (s *Service) ModifyBalance(mod *storetypes.ModifyBalance) (*storetypes.ModifyBalanceReturn, error) {
	bal, err := s.GetBalance(&storetypes.GetBalance{
		Account:  mod.Account,
		Currency: mod.Currency,
	})

	_ = bal.Balance.Add(bal.Balance, mod.Amount)

	_, err = s.db.Exec(`
		UPDATE
			balances
			LEFT JOIN accounts ON balances.account_id = accounts.id

		SET
			balances.balance = CAST($1 AS numeric)

		WHERE
			UPPER(accounts.id) = UPPER($2) AND
			UPPER(balances.currency) = UPPER($3)
	`, bal.Balance.String(), mod.Account, mod.Currency.String())

	return &storetypes.ModifyBalanceReturn{
		Account:  mod.Account,
		Currency: mod.Currency,
		Balance:  bal.Balance,
	}, err
}

func (s *Service) InsertOrder(order *storetypes.InsertOrder) (uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO
			orderbook(account_id, symbol, type, rate, quantity)

		VALUES
			(SELECT id FROM accounts WHERE UPPER(address) = UPPER($1)), $2, $3, $4, $5)

		RETURNING
			id;
	`, order.Account, order.Symbol.String(), order.Type.String(), order.Rate, order.Amount).Scan(&id); err != nil {
		log.Printf("err inserting order\n%v", err)
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateOrder(order *storetypes.UpdateOrder) error {
	_, err := s.db.Exec(`
		UPDATE
			orderbook
			LEFT JOIN accounts ON orderbook.account_id = accounts.id

		SET
			orderbook.symbol = $2,
			orderbook.type = $3,
			orderbook.rate = $4,
			orderbook.amount = $5

		WHERE
			orderbook.id = $6 AND
			accounts.address = $1
	`, order.Account, order.Symbol, order.Type, order.Rate, order.Amount, order.ID)

	return err
}

func (s *Service) CancelOrder(order *storetypes.CancelOrder) error {
	_, err := s.db.Exec(`
		DELETE
			orderbook

		FROM
			orderbook
			LEFT JOIN accounts ON orderbook.account_id = accounts.id

		WHERE
			orderbook.id = $1 AND
			UPPER(accounts.address) = UPPER($2) AND
			UPPER(orderbook.symbol) = UPPER($3)
	`, order.ID, order.Account, order.Symbol)

	return err
}

func (s *Service) GetOpenOrders(open *storetypes.GetOpenOrders) ([]*storetypes.GetOpenOrdersReturn, error) {
	var ret []*storetypes.GetOpenOrdersReturn

	threshArrow := "<"
	order := "DESC"
	if open.Type == ordertype.ASK {
		threshArrow = ">"
		order = "ASC"
	}

	rows, err := s.db.Query(fmt.Sprintf(`
		SELECT
			orderbook.id,
			accounts.address,
			orderbook.symbol,
			orderbook.type,
			orderbook.rate,
			orderbook.quantity::text

		FROM
			orderbook
			LEFT JOIN accounts ON orderbook.account_id = accounts.id

		WHERE
			UPPER(orderbook.type) = UPPER($1) AND
			orderbook.rate %s $2 AND
			UPPER(orderbook.symbol) = UPPER($3)

		ORDER BY
			orderbook.rate %s
	`, threshArrow, order), open.Type, open.Threshold, open.Symbol)
	if err != nil {
		if err == sql.ErrNoRows {
			return ret, nil
		}

		log.Printf("err querying\n%v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			accountAddress, s, t, quantity string
			id                             uint64
			rate                           float64
		)

		if err = rows.Scan(&id, &accountAddress, &s, &t, &rate, &quantity); err != nil {
			log.Printf("rows scan err\n%v", err)
			return nil, err
		}

		i := new(big.Int)
		_, ok := i.SetString(quantity, 10)
		if !ok {
			log.Println("err setting big int")
			return nil, errors.New("could not set big int")
		}

		symbol, err := symbol.SupportedTypeFromString(s)
		if err != nil {
			log.Printf("could not get symbol from stringn%v", err)
			return nil, err
		}
		typ, err := ordertype.SupportedOrderTypeFromString(t)
		if err != nil {
			log.Printf("could not get type from stringn%v", err)
			return nil, err
		}

		ret = append(ret, &storetypes.GetOpenOrdersReturn{
			ID:      id,
			Account: accountAddress,
			Symbol:  symbol,
			Type:    typ,
			Rate:    rate,
			Amount:  i,
		})
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows err\n%v", err)
		return nil, err
	}

	return ret, nil
}
