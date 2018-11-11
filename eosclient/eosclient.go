package eosclient

import (
	"log"
	"math/big"

	eos "github.com/eoscanada/eos-go"
)

// Config ...
type Config struct {
	URL   string
	Debug bool
}

// Client ...
type Client struct {
	client *eos.API
}

// NewClient ...
func NewClient(config *Config) *Client {
	client := eos.New(config.URL)

	if config.Debug {
		client.Debug = true
	}

	return &Client{
		client: client,
	}
}

// Info ...
func (s *Client) Info() (*eos.InfoResp, error) {
	return s.client.GetInfo()
}

// AccountInfo ...
func (s *Client) AccountInfo(account string) (*eos.AccountResp, error) {
	acct := eos.AccountName(account)
	return s.client.GetAccount(acct)
}

// SetSigner ...
func (s *Client) SetSigner(wifPrivateKey string) error {
	keyBag := eos.NewKeyBag()
	err := keyBag.ImportPrivateKey(wifPrivateKey)
	if err != nil {
		return err
	}

	s.client.SetSigner(keyBag)
	return nil
}

// Action ...
type Action struct {
	AccountName string
	ActionName  string
	Permissions string
	ActionData  interface{}
}

// PushAction ...
func (s *Client) PushAction(action *Action) (*eos.PushTransactionFullResp, error) {
	data := eos.ActionData{
		Data: action.ActionData,
	}

	perm, err := eos.NewPermissionLevel(action.Permissions)
	if err != nil {
		return nil, err
	}

	permissions := []eos.PermissionLevel{
		perm,
	}

	eosAction := &eos.Action{
		Account:       eos.AccountName(action.AccountName),
		Name:          eos.ActionName(action.ActionName),
		Authorization: permissions,
		ActionData:    data,
	}

	return s.client.SignPushActions(eosAction)
}

// GetTransaction ...
func (s *Client) GetTransaction(txID string) (*eos.TransactionResp, error) {
	return s.client.GetTransaction(txID)
}

// GetActions ...
func (s *Client) GetActions(accountName string, pos int, offset int) (*eos.ActionsResp, error) {
	return s.client.GetActions(eos.GetActionsRequest{
		AccountName: eos.AccountName(accountName),
		Pos:         int64(pos),
		Offset:      int64(offset),
	})
}

// OrderEvent ...
type OrderEvent struct {
	Sender string
	Price  *big.Int
	Amount *big.Int
	Value  *big.Int
}

// GetOrderEvents ...
func (s *Client) GetOrderEvents(accountName string) (chan *OrderEvent, error) {
	resp, err := s.GetActions(accountName, 0, 1000)
	if err != nil {
		return nil, err
	}

	results := make(chan *OrderEvent)

	go func() {
		for i := range resp.Actions {
			action := resp.Actions[i]

			if action.Trace.Action.Name == "placeorder" {
				data, ok := action.Trace.Action.Data.(map[string]interface{})
				if !ok {
					log.Fatal("not ok")
				}

				amountStr, ok := data["amount"].(string)
				if !ok {
					log.Fatal("not ok")
				}
				priceStr, ok := data["price"].(string)
				if !ok {
					log.Fatal("not ok")
				}
				valueStr, ok := data["value"].(string)
				if !ok {
					log.Fatal("not ok")
				}

				amount := new(big.Int)
				amount.SetString(amountStr, 10)

				price := new(big.Int)
				price.SetString(priceStr, 10)

				value := new(big.Int)
				value.SetString(valueStr, 10)

				result := &OrderEvent{
					Amount: amount,
					Price:  price,
					Value:  value,
				}

				results <- result
			}
		}
	}()

	return results, nil
}

// Withdraw ...
func (s *Client) Withdraw(receiver string, value *big.Int) (*eos.PushTransactionFullResp, error) {
	action := &Action{
		ActionName:  "withdraw",
		AccountName: "helloworld54",
		Permissions: "helloworld54@active",
		ActionData: struct {
			Receiver string `json:"receiver"`
			Value    uint64 `json:"value"`
		}{
			Receiver: receiver,
			Value:    value.Uint64(),
		},
	}

	resp, err := s.PushAction(action)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
