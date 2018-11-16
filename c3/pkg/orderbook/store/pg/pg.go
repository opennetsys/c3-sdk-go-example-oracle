package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"

	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/ordertype"
	storetypes "github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/store"
	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/symbol"
)

const (
	retries int           = 5
	wait    time.Duration = 5 * time.Second
)

func New(opts *Options) (*Service, error) {
	db, err := sql.Open("postgres", opts.PostgresURL)
	if err != nil {
		log.Printf("err opening postgres; err: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Printf("err pinging the db\n%v", err)
		if err = retry(db); err != nil {
			return nil, err
		}
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("err creating driver\n%v", err)
		return nil, err
	}

	gPath := os.Getenv("GOPATH")
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/src/github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/store/pg/migrations", gPath),
		"postgres", driver,
	)
	if err != nil {
		log.Printf("err creating migration\n%v", err)
		return nil, err
	}
	if err = m.Up(); err != nil && err.Error() != "no change" {
		log.Printf("err migrating\n%v", err)
		return nil, err
	}

	return &Service{
		opts: opts,
		db:   db,
	}, db.Ping()
}

func retry(db *sql.DB) error {
	time.Sleep(wait)
	for i := 0; i < retries; i++ {
		if err := db.Ping(); err != nil {
			time.Sleep(wait)
		} else {
			return nil
		}
	}

	return errors.New("couldn't connect to db")
}

func (s *Service) UpsertAccount(account *storetypes.UpsertAccount) (uint64, error) {
	var id uint64
	err := s.db.QueryRow(`
		INSERT INTO
			accounts(address, chain)

		VALUES
			($1, $2)

		ON CONFLICT("address", "chain")
			DO UPDATE SET chain = EXCLUDED.chain RETURNING id;
	`, account.Account, account.Chain.String()).Scan(&id)

	return id, err
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
			UPPER(accounts.address) = UPPER($1) AND
			UPPER(balances.currency) = UPPER($2)

		LIMIT
			1;
	`, get.Account, get.Currency.String()).Scan(&balance); err != nil {
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
	log.Printf("get %s", *get)
	fullBalance, err := s.GetBalance(get)
	if err != nil {
		log.Printf("err getting balance\n%v", err)
		return nil, err
	}
	log.Printf("full balance %s", fullBalance.Balance.String())

	// note: balance has non-null constraint
	rows, err := s.db.Query(`
		SELECT
			UPPER(symbol),
			type,
			rate,
			quantity::text

		FROM
			orderbook
			LEFT JOIN accounts ON orderbook.account_id = accounts.id

		WHERE
			UPPER(accounts.address) = UPPER($1)
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
	// note: this is bad but I'm sick of fighting SQL
	var count int
	newBal := new(big.Int)
	if err := s.db.QueryRow(`
		SELECT
			COUNT(*)

		FROM
			balances
			LEFT JOIN accounts ON balances.account_id = accounts.id

		WHERE
			UPPER(accounts.address) = UPPER($1) AND
			UPPER(balances.currency) = UPPER($2)
	`, mod.Account, mod.Currency.String()).Scan(&count); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("err counting\n%v", err)
			return nil, err
		}
	}
	if count == 0 {
		if _, err := s.db.Exec(`
			INSERT INTO
				balances(account_id, currency, balance)

			VALUES
				((SELECT id FROM accounts WHERE UPPER(address) = UPPER($2)), $3, CAST($1 AS numeric))
		`, mod.Amount.String(), mod.Account, mod.Currency.String()); err != nil {
			log.Printf("err inserting balance\n%v", err)
			return nil, err
		}
		newBal = mod.Amount

	} else {
		bal, err := s.GetBalance(&storetypes.GetBalance{
			Account:  mod.Account,
			Currency: mod.Currency,
		})
		if err != nil {
			log.Printf("err getting balance\n%v", err)
			return nil, err
		}

		_ = bal.Balance.Add(bal.Balance, mod.Amount)

		if _, err = s.db.Exec(`
			UPDATE
				balances

			SET
				balance = CAST($1 AS numeric)

			FROM
				accounts

			WHERE
				balances.account_id = accounts.id AND
				UPPER(accounts.address) = UPPER($2) AND
				UPPER(currency) = UPPER($3)
		`, bal.Balance.String(), mod.Account, mod.Currency.String()); err != nil {
			log.Printf("err updating balance\n%v", err)
			return nil, err
		}

		newBal = bal.Balance
	}

	return &storetypes.ModifyBalanceReturn{
		Account:  mod.Account,
		Currency: mod.Currency,
		Balance:  newBal,
	}, nil
}

func (s *Service) InsertOrder(order *storetypes.InsertOrder) (uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO
			orderbook(account_id, symbol, type, rate, quantity)

		VALUES
			((SELECT id FROM accounts WHERE UPPER(address) = UPPER($1)), $2, $3, $4, $5)

		RETURNING
			id;
	`, order.Account, order.Symbol.String(), order.Type.String(), order.Rate, order.Amount.String()).Scan(&id); err != nil {
		log.Printf("err inserting order\n%v", err)
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateOrder(order *storetypes.UpdateOrder) error {
	_, err := s.db.Exec(`
		UPDATE
			orderbook

		SET
			symbol = $2,
			type::text = $3,
			rate = $4,
			amount = $5

		FROM
			accounts 

		WHERE
			orderbook.account_id = accounts.id AND
			orderbook.id = $6 AND
			accounts.address = $1
	`, order.Account, order.Symbol.String(), order.Type.String(), order.Rate, order.Amount.String(), order.ID)

	return err
}

func (s *Service) CancelOrder(order *storetypes.CancelOrder) error {
	_, err := s.db.Exec(`
		DELETE FROM
			orderbook

		USING
			orderbook o1
			LEFT JOIN accounts ON o1.account_id = accounts.id

		WHERE
			o1.id = $1 AND
			UPPER(accounts.address) = UPPER($2) AND
			UPPER(o1.symbol) = UPPER($3)
	`, order.ID, order.Account, order.Symbol.String())

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
		orderbook.type::text = UPPER($1) AND
			orderbook.rate %s $2 AND
			UPPER(orderbook.symbol) = UPPER($3)

		ORDER BY
			orderbook.rate %s
	`, threshArrow, order), open.Type.String(), open.Threshold, open.Symbol.String())
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
