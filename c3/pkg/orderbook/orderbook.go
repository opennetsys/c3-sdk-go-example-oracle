package orderbook

import (
	"errors"
	"log"
	"math/big"
	"strings"

	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/currency"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/ordertype"
	storetypes "github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/store"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/store/pg"
)

func New(opts *Options) (*Service, error) {
	pgstore, err := pg.New(&pg.Options{
		PostgresURL: opts.PostgresURL,
	})

	// note: we don't check for err, we just return it
	return &Service{
		opts:  opts,
		store: pgstore,
	}, err
}

func (s *Service) Deposit(dep *Deposit) (*DepositReturn, error) {
	ret, err := s.store.ModifyBalance(&storetypes.ModifyBalance{
		Account:  dep.Account,
		Currency: dep.Currency,
		Amount:   dep.Amount,
	})

	return &DepositReturn{
		Account:  ret.Account,
		Currency: ret.Currency,
		Balance:  ret.Balance,
	}, err
}

func (s *Service) Withdrawal(with *Withdrawal) (*WithdrawalReturn, error) {
	bal, err := s.store.GetTradableBalance(&storetypes.GetBalance{
		Account:  with.Account,
		Currency: with.Currency,
	})
	if err != nil {
		log.Printf("err getting tradeable balane\n%v", err)
		return nil, err
	}

	if bal.Balance.Cmp(with.Amount) < 0 {
		return nil, errors.New("not enough funds")
	}

	z := new(big.Int)
	_ = z.Sub(with.Amount, bal.Balance)

	remainingBal, err := s.store.ModifyBalance(&storetypes.ModifyBalance{
		Account:  with.Account,
		Currency: with.Currency,
		Amount:   z,
	})
	if err != nil {
		log.Printf("err modifying balance\n%v", err)
		return nil, err
	}

	return &WithdrawalReturn{
		Account:  with.Account,
		Currency: with.Currency,
		Balance:  remainingBal.Balance,
	}, nil
}

func (s *Service) PlaceOrder(order *PlaceOrder) (*PlaceOrderReturn, error) {
	var (
		c   currency.SupportedCoinTyper
		err error
	)

	coins := strings.Split(order.Symbol.String(), "_")
	if len(coins) != 2 {
		log.Printf("expected len(coins) == 2; received %d", len(coins))
		return nil, errors.New("len(coins) != 2")
	}

	t := ordertype.BID
	if order.Type == ordertype.BID {
		t = ordertype.ASK
	}

	if t == ordertype.BID {
		c, err = currency.SupportedTypeFromString(coins[1])
	} else {
		c, err = currency.SupportedTypeFromString(coins[0])
	}

	bal, err := s.store.GetTradableBalance(&storetypes.GetBalance{
		Account:  order.Account,
		Currency: c,
	})
	if err != nil {
		log.Printf("err getting tradeable balane\n%v", err)
		return nil, err
	}

	if bal.Balance.Cmp(order.Amount) < 0 {
		return nil, errors.New("not enough funds")
	}

	openOrders, err := s.store.GetOpenOrders(&storetypes.GetOpenOrders{
		Symbol:    order.Symbol,
		Threshold: order.Rate,
		Type:      t,
	})
	if err != nil {
		log.Printf("err getting open orders\n%v", err)
		return nil, err
	}

	remaining := new(big.Int)
	_ = remaining.Set(order.Amount)
	zero := big.NewInt(0)
	for _, openOrder := range openOrders {
		if remaining.Cmp(zero) == 0 {
			break
		}

		coins := strings.Split(openOrder.Symbol.String(), "_")
		if len(coins) != 2 {
			log.Printf("expected len(coins) == 2; received: %d", len(coins))
			return nil, errors.New("len(coins) != 2")
		}

		var (
			c1, c2 currency.SupportedCoinTyper
		)
		a1 := new(big.Int)
		a2 := new(big.Int)

		switch remaining.Cmp(openOrder.Amount) {
		case -1:
			z := new(big.Int)
			_ = z.Sub(openOrder.Amount, remaining)

			err := s.store.UpdateOrder(&storetypes.UpdateOrder{
				ID:      openOrder.ID,
				Account: openOrder.Account,
				Symbol:  openOrder.Symbol,
				Type:    openOrder.Type,
				Rate:    openOrder.Rate,
				Amount:  z,
			})
			if err != nil {
				log.Printf("err updating order\n%v", err)
				return nil, err
			}

			if openOrder.Type == ordertype.BID {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(remaining)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)

				_ = a2.Set(remaining)
				_ = a2.Neg(a2)
			} else {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(remaining)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)
				_ = a1.Neg(a1)

				_ = a2.Set(remaining)

			}
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c1,
				Amount:   a1,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a1.Neg(a1)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c1,
				Amount:   a1,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a2.Neg(a2)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = remaining.SetInt64(0)

		case 0:
			err := s.store.CancelOrder(&storetypes.CancelOrder{
				ID:      openOrder.ID,
				Account: openOrder.Account,
				Symbol:  openOrder.Symbol,
			})
			if err != nil {
				log.Printf("err canceling order\n%v", err)
				return nil, err
			}

			if openOrder.Type == ordertype.BID {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(remaining)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)

				_ = a2.Set(remaining)
				_ = a2.Neg(a2)
			} else {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(remaining)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)
				_ = a1.Neg(a1)

				_ = a2.Set(remaining)

			}

			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c1,
				Amount:   a1,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a1.Neg(a1)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c1,
				Amount:   a1,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a2.Neg(a2)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = remaining.SetInt64(0)

		case 1:
			z := new(big.Int)
			_ = z.Sub(remaining, openOrder.Amount)

			err := s.store.CancelOrder(&storetypes.CancelOrder{
				ID:      openOrder.ID,
				Account: openOrder.Account,
				Symbol:  openOrder.Symbol,
			})
			if err != nil {
				log.Printf("err cancelling order\n%v", err)
				return nil, err
			}

			if openOrder.Type == ordertype.BID {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(openOrder.Amount)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)

				_ = a2.Set(openOrder.Amount)
				_ = a2.Neg(a2)
			} else {
				c1, err = currency.SupportedTypeFromString(coins[0])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}
				c2, err = currency.SupportedTypeFromString(coins[1])
				if err != nil {
					log.Printf("err getting coin from string\n%v", err)
					return nil, err
				}

				tmpRate := big.NewFloat(openOrder.Rate)
				tmpA1 := new(big.Float)
				tmpA1.SetInt(openOrder.Amount)
				_ = tmpA1.Quo(tmpA1, tmpRate)
				_, _ = tmpA1.Int(a1)
				_ = a1.Neg(a1)

				_ = a2.Set(openOrder.Amount)

			}

			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c1,
				Amount:   a1,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  openOrder.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a1.Neg(a1)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c1,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = a2.Neg(a2)
			_, err = s.store.ModifyBalance(&storetypes.ModifyBalance{
				Account:  order.Account,
				Currency: c2,
				Amount:   a2,
			})
			if err != nil {
				log.Printf("err updating balance\n%v", err)
				return nil, err
			}

			_ = remaining.Set(z)

		default:
			return nil, errors.New("unknown compare value")

		}
	}

	if remaining.Cmp(zero) > 0 {
		id, err := s.store.InsertOrder(&storetypes.InsertOrder{
			Account: order.Account,
			Symbol:  order.Symbol,
			Type:    order.Type,
			Rate:    order.Rate,
			Amount:  remaining,
		})
		if err != nil {
			log.Printf("err inserting order\n%v", err)
			return nil, err
		}

		_ = order.Amount.Sub(order.Amount, remaining)
		return &PlaceOrderReturn{
			ID:      &id,
			Account: order.Account,
			Symbol:  order.Symbol,
			Type:    order.Type,
			Filled:  order.Amount,
		}, nil
	}

	_ = order.Amount.Sub(order.Amount, remaining)
	return &PlaceOrderReturn{
		ID:      nil,
		Account: order.Account,
		Symbol:  order.Symbol,
		Type:    order.Type,
		Filled:  order.Amount,
	}, nil
}

func (s *Service) CancelOrder(order *CancelOrder) error {
	return s.store.CancelOrder(&storetypes.CancelOrder{
		ID:      order.ID,
		Account: order.Account,
		Symbol:  order.Symbol,
	})
}
