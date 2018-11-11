package orderbook

import (
	"math/big"

	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/chain"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/currency"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/ordertype"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/store"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/symbol"
)

type Options struct {
	PostgresURL string
}

type UpsertAccount struct {
	Account string                    `json:"account"`
	Chain   chain.SupportedChainTyper `json:"chain"`
}

type Service struct {
	opts  *Options
	store store.Interface
}

type Deposit struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Amount   *big.Int                    `json:"amount"`
}

type DepositReturn struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Balance  *big.Int                    `json:"balance"`
}

type Withdrawal struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Amount   *big.Int                    `json:"amount"`
}

type WithdrawalReturn struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Balance  *big.Int                    `json:"balance"`
}

type PlaceOrder struct {
	Account string                            `json:"account"`
	Symbol  symbol.SupportedSymbolTyper       `json:"symbol"`
	Type    ordertype.SupportedOrderTypeTyper `json:"orderType"`
	Rate    float64                           `json:"rate"`
	Amount  *big.Int                          `json:"amount"`
}

type PlaceOrderReturn struct {
	ID      *uint64                           `json:"id"`
	Account string                            `json:"account"`
	Symbol  symbol.SupportedSymbolTyper       `json:"symbol"`
	Type    ordertype.SupportedOrderTypeTyper `json:"orderType"`
	Filled  *big.Int                          `json:"filled"`
}

type CancelOrder struct {
	ID      uint64                      `json:"id"`
	Account string                      `json:"account"`
	Symbol  symbol.SupportedSymbolTyper `json:"symbol"`
}
