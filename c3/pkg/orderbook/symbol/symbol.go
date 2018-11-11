package symbol

import (
	"errors"
	"strings"
)

type supported int

const (
	EOS_ETH supported = iota
)

// SupportedSymbolTyper are the coins supported by the packer
type SupportedSymbolTyper interface {
	Type() supported
	String() string
}

// every base must fullfill the supported interface
func (s supported) Type() supported {
	return s
}

// AllSupportedSymbols returns all of the supported coins
func AllSupportedSymbols() []SupportedSymbolTyper {
	return []SupportedSymbolTyper{
		EOS_ETH,
	}
}

// SupportedTypeFromString parses a string to return the supported type
func SupportedTypeFromString(s string) (SupportedSymbolTyper, error) {
	switch strings.ToUpper(s) {
	case "EOS_ETH":
		return EOS_ETH, nil

	default:
		return nil, errors.New("unsupported symbol type")
	}
}
