#!/bin/bash

service postgresql restart &&\
  go run /go/src/github.com/c3systems/Hackathon-EOS-SF-2018/cmd/dapp/main.go
