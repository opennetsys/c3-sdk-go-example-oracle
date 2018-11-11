package chain

import (
	"errors"
	"strings"
)

type supported int

const (
	ETHEREUM supported = iota
	EOSIO
)

// SupportedChainTyper are the chains supported by the packer
type SupportedChainTyper interface {
	Type() supported
	String() string
}

// every base must fullfill the supported interface
func (s supported) Type() supported {
	return s
}

// AllSupportedChain returns all of the supported chains
func AllSupportedChains() []SupportedChainTyper {
	return []SupportedChainTyper{
		ETHEREUM,
		EOSIO,
	}
}

// SupportedChainFromString parses a string to return the supported type
func SupportedChainFromString(s string) (SupportedChainTyper, error) {
	switch strings.ToUpper(s) {
	case "ETHEREUM":
		return ETHEREUM, nil

	case "EOSIO":
		return EOSIO, nil

	default:
		return nil, errors.New("unsupported chain type")
	}
}
