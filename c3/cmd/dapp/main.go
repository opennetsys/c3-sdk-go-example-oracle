package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/cfg"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/coder"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/ethereumclient"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/currency"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/ordertype"
	"github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/orderbook/symbol"

	c3 "github.com/c3systems/c3-sdk-go"
)

var (
	client    = c3.NewC3()
	vars      cfg.Vars
	ethClient *ethereumclient.Client
	obook     *orderbook.Service
)

const (
	key = "data"
)

// App ...
type App struct {
}

func (s *App) processETHDeposit(depStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...
	dep, err := coder.DecodeETHLogDeposit(depStr)
	if err != nil {
		log.Printf("err decoding eth log\n%v", err)
		return err
	}

	receipt, err := obook.Deposit(&obook.Deposit{
		Account:  dep.Sender.GetHex(),
		Currency: currency.ETH,
		Amount:   dep.Value,
	})
	if err != nil {
		log.Printf("err depositing eth\n%v", err)
		return err
	}
	log.Printf("deposit receipt\n%v", receipt)

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}

func (s *App) processETHWithdrawal(withdrawalStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...
	withdrawal, err := coder.DecodeETHLogWithdrawal(withdrawalStr)
	if err != nil {
		log.Printf("err decoding eth log\n%v", err)
		return err
	}

	receipt, err := obook.Withdrawal(&obook.Withdrawal{
		Account:  withdrawal.Account.GetHex(),
		Currency: currency.ETH,
		Amount:   withdrawal.Value,
	})
	if err != nil {
		log.Printf("err withdrawing eth in c3\n%v", err)
		return err
	}
	log.Printf("c3 withdrawal receipt\n%v", err)

	res, err := ethClient.Withdraw(withdrawal.Account.GetHex(), withdrawal.Value)
	if err != nil {
		log.Printf("err withdrawing eth at the node\n%v", err)
		return err
	}
	log.Printf("ethereum withdrawal receipt\n%v", res)

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}
func (s *App) processETHBuy(buyStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...
	buy, err := coder.DecodeETHLogBuy(buyStr)
	if err != nil {
		log.Printf("err decoding eth log\n%v", err)
		return err
	}

	receipt, err := obook.PlaceOrder(&obook.PlaceOrder{
		Account: buy.Sender.GetHex(),
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.BID,              // note: is this correct?
		Rate:    float64(buy.Price.Int64()), // I'm assuming the rate was calucated correctly before this point
		Amount:  buy.Value,
	})
	if err != nil {
		log.Printf("err placing order\n%v", err)
		return err
	}
	log.Printf("c3 purchase receipt\n%v", receipt)

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}

func (s *App) processEOSDeposit(buyStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}

func (s *App) processEOSWithdrawal(withdrawalStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}
func (s *App) processEOSBuy(buyStr string) error {
	if err := setState(); err != nil {
		log.Printf("err setting state\n%v", err)
		return err
	}

	// do stuff here ...

	if err := getState(); err != nil {
		log.Printf("err getting state\n%v", err)
		return err
	}

	return nil
}

func startC3() {
	data := &App{}
	if err := client.RegisterMethod("processETHDeposit", []string{"string"}, data.processETHDeposit); err != nil {
		log.Fatal(err)
	}
	if err := client.RegisterMethod("processETHWithdrawal", []string{"string"}, data.processETHWithdrawal); err != nil {
		log.Fatal(err)
	}
	if err := client.RegisterMethod("processETHBuy", []string{"string"}, data.processETHBuy); err != nil {
		log.Fatal(err)
	}
	if err := client.RegisterMethod("processEOSDeposit", []string{"string"}, data.processEOSDeposit); err != nil {
		log.Fatal(err)
	}
	if err := client.RegisterMethod("processEOSWithdrawal", []string{"string"}, data.processEOSWithdrawal); err != nil {
		log.Fatal(err)
	}
	if err := client.RegisterMethod("processEOSBuy", []string{"string"}, data.processEOSBuy); err != nil {
		log.Fatal(err)
	}
	client.Serve()
}

func main() {
	// 1. run the cli
	constants, err := cfg.New(os.Args)
	if err != nil {
		log.Fatalf("Error grabbing cfg: %v", err)
	}
	vars = constants.Get()

	// 2. build the orderbook
	obook, err = orderbook.New(&obook.Options{
		PostgresURL: vars.PostgresURL,
	})
	if err != nil {
		log.Fatalf("err starting the orderbook\n%v", err)
	}

	// 3. build the eth client
	ch := make(chan interface{})
	ethClient, err = ethereumclient.NewClient(&ethereumclient{
		NodeURL:         vars.ETH_NodeURL,
		PrivateKey:      vars.ETH_PrivateKey,
		ContractAddress: vars.ETH_ContractAddress,
		ListenChan:      ch,
	})
	if err != nil {
		log.Fatalf("err building the eth client\n%v", err)
	}

	// 4. stat c3
	go startC3()

	// 5. wait
	select {}
}

func setState() error {
	prevState, found := client.State().Get([]byte(key))
	if !found {
		return errors.New("no previous state")
	}
	gPath := os.Getenv("GOPATH")

	if err := os.Remove(fmt.Sprintf("%s/src/github.com/c3systems/Hackathon-EOS-SF-2018/c3/cmd/dapp/state.tar", gPath)); err != nil {
		log.Printf("err removing prev state.tar\n%v", err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s/src/github.com/c3systems/Hackathon-EOS-SF-2018/c3/cmd/dapp/state.tar", gPath), prevState, 0644); err != nil {
		log.Printf("err writing state.tar\n%v", err)
		return err
	}

	cmd := exec.Command("/bin/sh", fmt.Sprintf("%s/src/github.com/c3systems/Hackathon-EOS-SF-2018/c3/cmd/dapp/set_state.sh", gPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("err running set_state\n%v\nnoutput:\n%s", err, string(out))
		return err
	}

}

func getState() error {
	gPath := os.Getenv("GOPATH")
	cmd := exec.Command("/bin/sh", fmt.Sprintf("%s/src/github.com/c3systems/Hackathon-EOS-SF-2018/c3/cmd/dapp/get_state.sh", gPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("err running get_state\n%v\nnoutput:\n%s", err, string(out))
		return err
	}

	stateBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/c3systems/Hackathon-EOS-SF-2018/c3/cmd/dapp/state.tar", gPath))
	if err != nil {
		log.Printf("err reading state tar file\n%v", err)
		return err
	}

	return client.State().Set([]byte(key), stateBytes)
}
