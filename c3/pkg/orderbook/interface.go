package orderbook

type Interface interface {
	UpsertAccount(account *UpsertAccount) (uint64, error)
	Deposit(dep *Deposit) (*DepositReturn, error)
	Withdrawal(with *Withdrawal) (*WithdrawalReturn, error)
	PlaceOrder(order *PlaceOrder) (*PlaceOrderReturn, error)
	CancelOrder(order *CancelOrder) error
}
