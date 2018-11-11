package store

type Interface interface {
	GetBalance(get *GetBalance) (*GetBalanceReturn, error)
	GetTradableBalance(get *GetBalance) (*GetBalanceReturn, error)
	ModifyBalance(mod *ModifyBalance) (*ModifyBalanceReturn, error)
	InsertOrder(order *InsertOrder) (uint64, error)
	UpdateOrder(order *UpdateOrder) error
	CancelOrder(order *CancelOrder) error
	GetOpenOrders(open *GetOpenOrders) ([]*GetOpenOrdersReturn, error)
}
