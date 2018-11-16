package coder

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/ethereumclient"
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
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	// Decode (receive) the value.
	var l ethereumclient.LogBuy
	if err := dec.Decode(&l); err != nil {
		log.Printf("decode error:\n%v", err)
		return nil, err
	}

	return &l, nil
}
func DecodeETHLogWithdrawal(b []byte) (*ethereumclient.LogWithdrawal, error) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	// Decode (receive) the value.
	var l ethereumclient.LogWithdrawal
	if err := dec.Decode(&l); err != nil {
		log.Printf("decode error:\n%v", err)
		return nil, err
	}

	return &l, nil
}
func DecodeETHLogDeposit(b []byte) (*ethereumclient.LogDeposit, error) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	// Decode (receive) the value.
	var l ethereumclient.LogDeposit
	if err := dec.Decode(&l); err != nil {
		log.Printf("decode error:\n%v", err)
		return nil, err
	}

	return &l, nil
}
