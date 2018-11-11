package currency

import (
	"errors"
	"strings"
)

type supported int

const (
	ETH supported = iota
	EOS
)

// SupportedCoinTyper are the coins supported by the packer
type SupportedCoinTyper interface {
	Type() supported
	String() string
}

// every base must fullfill the supported interface
func (s supported) Type() supported {
	return s
}

// AllSupportedCoins returns all of the supported coins
func AllSupportedCoins() []SupportedCoinTyper {
	return []SupportedCoinTyper{
		ETH,
		EOS,
	}
}

// SupportedTypeFromString parses a string to return the supported type
func SupportedTypeFromString(s string) (SupportedCoinTyper, error) {
	switch strings.ToUpper(s) {
	case "ETH":
		return ETH, nil

	case "EOS":
		return EOS, nil

	default:
		return nil, errors.New("unsupported coint type")
	}
}
