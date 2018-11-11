package ethereumclient

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	auth     *bind.TransactOpts
	client   *ethclient.Client
	instance *contract.Exchange
}

// NewClient ...
func NewClient(config *Config) *Client {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(config.ContractAddress)
	instance, err := contract.NewExchange(address, client)
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
		auth:     auth,
		client:   client,
		instance: instance,
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
