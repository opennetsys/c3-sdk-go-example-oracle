package orderbook

type Interface interface {
	Deposit(dep *Deposit) (*DepositReturn, error)
	Withdrawal(with *Withdrawal) (*WithdrawalReturn, error)
	Trade(t *Trade) (*TradeReturn, error)
	CancelTrade(c *CancelTrade) (*CancelTradeReturn, error)
}
