package coder

import (
	"bytes"
	"encoding/gob"
	"log"
)

func EncodeETHLogBuy(l *ethereumclient.LogBuy) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(*l)
	if err != nil {
		log.Printf("encode err\n%v", err)
		return nil, err
	}

	return network.Bytes(), nil
}
func EncodeETHLogWithdrawal(l *ethereumclient.LogWithdrawal) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(*l)
	if err != nil {
		log.Printf("encode err\n%v", err)
		return nil, err
	}

	return network.Bytes(), nil
}
func EncodeETHLogDeposit(l *ethereumclient.LogDeposit) ([]byte, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(*l)
	if err != nil {
		log.Printf("encode err\n%v", err)
		return nil, err
	}

	return network.Bytes(), nil
}
func DecodeETHLogBuy(b []byte) (*ethereumclient.LogBuy, error) {
	dec := gob.NewDecoder(b)

	// Decode (receive) the value.
	var l ethereumclient.LogBuy
	err = dec.Decode(&l)
	if err != nil {
		log.Printf("decode error:", err)
		return nil, err
	}

	return &l, nil
}
func DecodeETHLogWithdrawal(b []byte) (*ethereumclient.LogWithdrawal, error) {
	dec := gob.NewDecoder(b)

	// Decode (receive) the value.
	var l ethereumclient.LogWithdrawal
	err = dec.Decode(&l)
	if err != nil {
		log.Printf("decode error:", err)
		return nil, err
	}

	return &l, nil
}
func DecodeETHLogDeposit(b []byte) (*ethereumclient.LogDeposit, error) {
	dec := gob.NewDecoder(b)

	// Decode (receive) the value.
	var l ethereumclient.LogDeposit
	err = dec.Decode(&l)
	if err != nil {
		log.Printf("decode error:", err)
		return nil, err
	}

	return &l, nil
}
