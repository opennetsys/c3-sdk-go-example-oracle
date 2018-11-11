package ethereumclient

import (
	"context"
	"crypto/ecdsa"
	"errors"
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

	contract "github.com/c3systems/Hackathon-EOS-SF-2018/c3/pkg/contracts"
)

// Config ...
type Config struct {
	NodeURL         string
	PrivateKey      string
	ContractAddress string
	ListenChan      chan interface{}
}

// Client ...
type Client struct {
	auth            *bind.TransactOpts
	client          *ethclient.Client
	instance        *contract.Exchange
	contractAddress common.Address
	listenChan      chan interface{}
}

// NewClient ...
func NewClient(config *Config) (*Client, error) {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := contract.NewExchange(contractAddress, client)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
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
	}, nil
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
	Price  *big.Int // note: should probably be a float? so in solidity that would be a numerator and denom
	Value  *big.Int
}

// LogDeposit ...
type LogDeposit struct {
	Sender common.Address
	Value  *big.Int
}

// LogWithdrawal ...
type LogWithdrawal struct {
	Receiver common.Address
	Value    *big.Int
}

// Listen ...
func (s *Client) ListenBuy() {
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
			s.listenChan <- err

		case vLog := <-logs:
			fmt.Print("HERE")
			switch vLog.Topics[0].Hex() {
			case logBuySigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var event LogBuy

				err := contractAbi.Unpack(&event, "LogBuy", vLog.Data)
				if err != nil {
					log.Printf("err unpacking abi\n%v", err)
					s.listenChan <- err
					continue
				}

				fmt.Printf("Event: %s\n", event.Sender)
				s.listenChan <- &event

			default:
				log.Printf("unknown vlog topic\n%v", vLog)
				s.listenChan <- errors.New("unknown vLog Topics type")

			}
		}
	}
}
func (s *Client) ListenDeposit() {
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

	logDepositSig := []byte("LogDeposit(address,uint256)")
	logDepositSigHash := crypto.Keccak256Hash(logDepositSig)

	for {
		select {
		case err := <-sub.Err():
			s.listenChan <- err

		case vLog := <-logs:
			fmt.Print("HERE - DEPOSIT")
			switch vLog.Topics[0].Hex() {
			case logDepositSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var event LogDeposit

				err := contractAbi.Unpack(&event, "LogDeposit", vLog.Data)
				if err != nil {
					log.Printf("err unpacking abi\n%v", err)
					s.listenChan <- err
					continue
				}

				fmt.Printf("Event: %s\n", event.Sender)
				s.listenChan <- &event

			default:
				log.Printf("unknown vlog topic type %T\n%v", vLog, vLog)
				s.listenChan <- errors.New("unknown vLog Topics type")

			}
		}
	}
}

func (s *Client) ListenWithdrawal() {
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

	logWithdrawalSig := []byte("LogWithdrawal(address,uint256)")
	logWithdrawalSigHash := crypto.Keccak256Hash(logWithdrawalSig)

	for {
		select {
		case err := <-sub.Err():
			s.listenChan <- err

		case vLog := <-logs:
			fmt.Print("HERE")
			switch vLog.Topics[0].Hex() {
			case logWithdrawalSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var event LogWithdrawal

				err := contractAbi.Unpack(&event, "LogWithdrawal", vLog.Data)
				if err != nil {
					log.Printf("err unpacking abi\n%v", err)
					s.listenChan <- err
					continue
				}

				fmt.Printf("Event: %s\n", event.Receiver)
				s.listenChan <- &event

			default:
				log.Printf("unknown vlog topic\n%v", vLog)
				s.listenChan <- errors.New("unknown vLog Topics type")

			}
		}
	}
}
