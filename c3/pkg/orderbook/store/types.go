package store

import (
	"math/big"

	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/chain"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/currency"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/ordertype"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/symbol"
)

type UpsertAccount struct {
	Account string                    `json:"account"`
	Chain   chain.SupportedChainTyper `json:"chain"`
}

type ModifyBalance struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Amount   *big.Int                    `json:"amount"`
}

type ModifyBalanceReturn struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Balance  *big.Int                    `json:"balance"`
}

type GetBalance struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
}

type GetBalanceReturn struct {
	Account  string                      `json:"account"`
	Currency currency.SupportedCoinTyper `json:"currency"`
	Balance  *big.Int                    `json:"balance"`
}

type InsertOrder struct {
	Account string                            `json:"account"`
	Symbol  symbol.SupportedSymbolTyper       `json:"symbol"`
	Type    ordertype.SupportedOrderTypeTyper `json:"orderType"`
	Rate    float64                           `json:"rate"`
	Amount  *big.Int                          `json:"amount"`
}

type GetOpenOrders struct {
	Symbol    symbol.SupportedSymbolTyper       `json:"symbol"`
	Threshold float64                           `json:"amount"`
	Type      ordertype.SupportedOrderTypeTyper `json:"orderType"`
}

type GetOpenOrdersReturn struct {
	ID      uint64                            `json:"id"`
	Account string                            `json:"account"`
	Symbol  symbol.SupportedSymbolTyper       `json:"symbol"`
	Type    ordertype.SupportedOrderTypeTyper `json:"orderType"`
	Rate    float64                           `json:"rate"`
	Amount  *big.Int                          `json:"amount"`
}

type UpdateOrder struct {
	ID      uint64                            `json:"id"`
	Account string                            `json:"account"`
	Symbol  symbol.SupportedSymbolTyper       `json:"symbol"`
	Type    ordertype.SupportedOrderTypeTyper `json:"orderType"`
	Rate    float64                           `json:"rate"`
	Amount  *big.Int                          `json:"amount"`
}

type CancelOrder struct {
	ID      uint64                      `json:"id"`
	Account string                      `json:"account"`
	Symbol  symbol.SupportedSymbolTyper `json:"symbol"`
}
