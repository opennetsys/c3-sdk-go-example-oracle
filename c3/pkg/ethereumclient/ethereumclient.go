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

	contract "github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/contracts"
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
		log.Printf("err dialing\n%v", err)
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := contract.NewExchange(contractAddress, client)
	if err != nil {
		log.Printf("err creating instance\n%v", err)
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Printf("err converting private key hex to ecdsa\n%v", err)
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
		log.Printf("err getting nonce\n%v", err)
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("err getting suggested gas price\n%v", err)
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
		listenChan:      config.ListenChan,
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
func (s *Client) Listen() {
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

	logDepositSig := []byte("LogDeposit(address,uint256)")
	logDepositSigHash := crypto.Keccak256Hash(logDepositSig)

	logBuySig := []byte("LogBuy(address,uint256,uint256,uint256)")
	logBuySigHash := crypto.Keccak256Hash(logBuySig)

	for {
		select {
		case err := <-sub.Err():
			s.listenChan <- err

		case vLog := <-logs:
			fmt.Print("HERE")
			switch strings.ToLower(vLog.Topics[0].Hex()) {
			case strings.ToLower(logBuySigHash.Hex()):
				fmt.Printf("Log Buy Name: Transfer\n")

				var event LogBuy

				err := contractAbi.Unpack(&event, "LogBuy", vLog.Data)
				if err != nil {
					log.Printf("err unpacking abi\n%v", err)
					s.listenChan <- err
					continue
				}

				fmt.Printf("Event: %s\n", event.Sender)
				s.listenChan <- &event

			case strings.ToLower(logDepositSigHash.Hex()):
				fmt.Printf("Log Deposit Name: Transfer\n")

				var event LogDeposit

				err := contractAbi.Unpack(&event, "LogDeposit", vLog.Data)
				if err != nil {
					log.Printf("err unpacking abi\n%v", err)
					s.listenChan <- err
					continue
				}

				fmt.Printf("Event: %s\n", event.Sender)
				s.listenChan <- &event

			case strings.ToLower(logWithdrawalSigHash.Hex()):
				fmt.Printf("Log Withdrawal Name: Transfer\n")

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
				log.Printf("unknown vlog topic\ntype %T\n%s\nexpected %s", vLog, vLog.Topics[0].Hex(), logWithdrawalSigHash.Hex())
				s.listenChan <- errors.New("unknown vLog Topics type")

			}
		}
	}
}
