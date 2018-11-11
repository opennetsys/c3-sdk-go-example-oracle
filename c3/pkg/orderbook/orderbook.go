package orderbook

import "github.com/c3systems/Hackathon-EOS-SF-2018/pkg/orderbook/store/pg"

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

}

func (s *Service) Withdrawal(with *Withdrawal) (*WithdrawalReturn, error) {

}

func (s *Service) Trade(t *Trade) (*TradeReturn, error) {

}

func (s *Service) CancelTrade(c *CancelTrade) (*CancelTradeReturn, error) {

}
