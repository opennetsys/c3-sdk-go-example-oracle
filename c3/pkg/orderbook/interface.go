package orderbook

type Interface interface {
	Deposit(dep *Deposit) (*DepositReturn, error)
	Withdrawal(with *Withdrawal) (*WithdrawalReturn, error)
	PlaceOrder(order *PlaceOrder) (*PlaceOrderReturn, error)
	CancelOrder(order *CancelOrder) error
}
