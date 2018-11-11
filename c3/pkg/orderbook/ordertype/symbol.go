package ordertype

import (
	"errors"
	"strings"
)

type supported int

const (
	BID supported = iota
	ASK
)

// SupportedOrderTypeTyper are the orders supported by the packer
type SupportedOrderTypeTyper interface {
	Type() supported
	String() string
}

// every base must fullfill the supported interface
func (s supported) Type() supported {
	return s
}

// AllSupportedOrderTypes returns all of the supported order types
func AllSupportedOrderTypes() []SupportedOrderTypeTyper {
	return []SupportedOrderTypeTyper{
		BID,
		ASK,
	}
}

// SupportedOrderTypeFromString parses a string to return the supported type
func SupportedOrderTypeFromString(s string) (SupportedOrderTypeTyper, error) {
	switch strings.ToUpper(s) {
	case "BID":
		return BID, nil

	case "ASK":
		return ASK, nil

	default:
		return nil, errors.New("unsupported order type")
	}
}
