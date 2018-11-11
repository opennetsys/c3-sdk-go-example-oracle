package ethereumclient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contract "../contracts" // TODO
)

// Config ...
type Config struct {
	NodeURL         string
	PrivateKey      string
	ContractAddress string
}

// Client ...
type Client struct {
	auth            *bind.TransactOpts
	client          *ethclient.Client
	instance        *contract.Exchange
	contractAddress common.Address
}

// NewClient ...
func NewClient(config *Config) *Client {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := contract.NewExchange(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return &Client{
		auth:            auth,
		client:          client,
		instance:        instance,
		contractAddress: contractAddress,
	}
}

// Withdraw ...
func (s *Client) Withdraw(receiver string, value *big.Int) (string, error) {
	tx, err := s.instance.Withdraw(s.auth, common.HexToAddress(receiver), value)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

// LogBuy ...
type LogBuy struct {
	Sender common.Address
	Amount *big.Int
	Price  *big.Int
	Value  *big.Int
}

// Listen ...
func (s *Client) Listen() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ExchangeABI)))
	if err != nil {
		log.Fatal(err)
	}

	logBuySig := []byte("LogBuy(address,uint256,uint256,uint256)")
	logBuySigHash := crypto.Keccak256Hash(logBuySig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Print("HERE")
			switch vLog.Topics[0].Hex() {
			case logBuySigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var event LogBuy

				err := contractAbi.Unpack(&event, "LogBuy", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Event: %s\n", event.Sender)
			}
		}
	}

	return nil
}
