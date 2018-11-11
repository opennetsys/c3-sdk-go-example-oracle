package orderbook

import (
	"github.com/c3systems/Hackathon-EOS-SF-2018/pkg/orderbook/store"
)

type Options struct {
	PostgresURL string
}

type Service struct {
	opts  *Options
	store store.Interface
}

type Deposit struct{}

type DepositReturn struct{}

type Withdrawal struct{}

type WithdrawalReturn struct{}

type Trade struct{}

type TradeReturn struct{}

type CancelTrade struct{}

type CancelTradeReturn struct{}
