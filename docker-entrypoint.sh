#!/bin/bash

/etc/init.d/postgresql start &&\
  sh ./wait-for-postgres.sh &&\
  psql -U postgres --command "CREATE DATABASE db;" &&\
  psql -U postgres --command "CREATE USER docker WITH SUPERUSER; ALTER USER docker VALID UNTIL 'infinity'; GRANT ALL PRIVILEGES ON DATABASE db TO docker;" &&\
  go run /go/src/github.com/c3systems/c3-sdk-go-example-oracle/c3/cmd/dapp/main.go
