package eosclient

import (
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
