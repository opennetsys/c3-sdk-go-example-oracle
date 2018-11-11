package eosclient

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Skip()
	_ = NewClient(&Config{
		URL: "http://api.kylin.alohaeos.com",
	})
}

func TestInfo(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL: "http://api.kylin.alohaeos.com",
	})

	info, err := client.Info()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(info)
}

func TestAccountInfo(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL: "http://api.kylin.alohaeos.com",
	})

	info, err := client.AccountInfo("helloworld54")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(info)
}

func TestGetTransaction(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL:   "https://api-kylin.eosasia.one",
		Debug: true,
	})

	tx, err := client.GetTransaction("3d43785ceca9a919e73b547487d9da6dad246f05425e513035e373c67310bc47")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tx)
}

func TestPushAction(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL:   "https://api-kylin.eosasia.one",
		Debug: true,
	})

	action := &Action{
		ActionName:  "placeorder",
		AccountName: "helloworld54",
		Permissions: "helloworld54@active",
		ActionData: struct {
			Sender string `json:"sender"`
			Price  uint64 `json:"price"`
			Amount uint64 `json:"amount"`
			Value  uint64 `json:"value"`
		}{
			Sender: "helloworld54",
			Price:  uint64(1),
			Amount: uint64(1),
			Value:  uint64(1),
		},
	}

	wifPrivateKey := "5Jh9tD4Fp1EpVn3EzEW6ura5NV3NddY8NNBcfpCZTvPDsKd9i5c"
	client.SetSigner(wifPrivateKey)

	resp, err := client.PushAction(action)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp)
}

func TestGetActions(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL:   "https://api-kylin.eosasia.one",
		Debug: true,
	})

	resp, err := client.GetActions("helloworld54", 0, 1000)
	if err != nil {
		t.Error(err)
	}

	_ = resp
}

func TestGetOrderEvents(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		URL:   "https://api-kylin.eosasia.one",
		Debug: true,
	})

	events, err := client.GetOrderEvents("helloworld54")
	if err != nil {
		t.Error(err)
	}

	event := <-events
	fmt.Println(event)
}

func TestWithdraw(t *testing.T) {
	client := NewClient(&Config{
		URL:   "https://api-kylin.eosasia.one",
		Debug: true,
	})

	wifPrivateKey := "5Jh9tD4Fp1EpVn3EzEW6ura5NV3NddY8NNBcfpCZTvPDsKd9i5c"
	client.SetSigner(wifPrivateKey)

	resp, err := client.Withdraw("myaccount123", big.NewInt(1))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp)
}
