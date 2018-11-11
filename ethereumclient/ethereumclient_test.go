package ethereumclient

import (
	"math/big"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Skip()
	client := NewClient(&Config{
		NodeURL:         "http://localhost:8545",
		PrivateKey:      "c98ebdb872cc52821e40a144ab636b8e072d151b3e6dda1ac4409a0014ea3155",
		ContractAddress: "0x197159a3fd77fff557206aec108516fd13e84189",
	})

	_ = client
}

func TestWithdraw(t *testing.T) {
	client := NewClient(&Config{
		NodeURL:         "http://localhost:8545",
		PrivateKey:      "c98ebdb872cc52821e40a144ab636b8e072d151b3e6dda1ac4409a0014ea3155",
		ContractAddress: "0xff40f9e5e3b392fd3bdc5990f007beda20d0290f",
	})

	receiver := "0x656f3db0b3a18a0e2c80f7d55f8eb9fd813e19c2"
	value := big.NewInt(1)
	tx, err := client.Withdraw(receiver, value)
	if err != nil {
		t.Error(err)
	}

	t.Log(tx)
}
