// +build integration

package orderbook

import (
	"database/sql"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/chain"
	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/currency"
	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/ordertype"
	"github.com/c3systems/c3-sdk-go-example-oracle/c3/pkg/orderbook/symbol"
	"github.com/joho/godotenv"
)

func TestPlaceOrder(t *testing.T) {
	// 1. connect to the db
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	TEST_POSTGRES_URL := os.Getenv("TEST_POSTGRES_URL")
	if TEST_POSTGRES_URL == "" {
		t.Fatalf("TEST_POSTGRES_URL env var is required")
	}

	db, err := sql.Open("postgres", TEST_POSTGRES_URL)
	if err != nil {
		t.Fatalf("err opening db: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("err pinging db: %v", err)
	}

	obook, err := New(&Options{
		PostgresURL: TEST_POSTGRES_URL,
	})
	if err != nil {
		t.Fatalf("err opening db: %v", err)
	}

	// 2. Insert some test data
	id1, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eth1",
		Chain:   chain.ETHEREUM,
	})
	_ = id1
	id2, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eth2",
		Chain:   chain.ETHEREUM,
	})
	_ = id2
	id3, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eth3",
		Chain:   chain.ETHEREUM,
	})
	_ = id3
	id4, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eos1",
		Chain:   chain.EOSIO,
	})
	_ = id4
	id5, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eos2",
		Chain:   chain.EOSIO,
	})
	_ = id5
	id6, err := obook.UpsertAccount(&UpsertAccount{
		Account: "eos3",
		Chain:   chain.EOSIO,
	})
	_ = id6

	_, err = obook.Deposit(&Deposit{
		Account:  "eth1",
		Currency: currency.ETH,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}
	_, err = obook.Deposit(&Deposit{
		Account:  "eth2",
		Currency: currency.ETH,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}
	_, err = obook.Deposit(&Deposit{
		Account:  "eth3",
		Currency: currency.ETH,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}
	_, err = obook.Deposit(&Deposit{
		Account:  "eos1",
		Currency: currency.EOS,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}
	_, err = obook.Deposit(&Deposit{
		Account:  "eos2",
		Currency: currency.EOS,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}
	_, err = obook.Deposit(&Deposit{
		Account:  "eos3",
		Currency: currency.EOS,
		Amount:   big.NewInt(10),
	})
	if err != nil {
		t.Fatalf("err depositing\n%v", err)
	}

	// defer cleanup
	defer func() {
		if _, err := db.Exec(`
			TRUNCATE
				accounts

			CASCADE;
		`); err != nil {
			log.Printf("err cleaning up\n%v", err)
		}
	}()

	// try an order that should fail for insufficient funds
	res, err := obook.PlaceOrder(&PlaceOrder{
		Account: "eth1",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.BID,
		Rate:    1,
		Amount:  big.NewInt(10000),
	})
	if err == nil {
		t.Fatal("expected first insufficient funds err")
	}

	// try to fail in another way
	res, err = obook.PlaceOrder(&PlaceOrder{
		Account: "eth1",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.ASK,
		Rate:    1,
		Amount:  big.NewInt(1),
	})
	if err == nil {
		t.Fatal("expected second insufficient funds err")
	}
	_ = res

	// succeed
	resO, err := obook.PlaceOrder(&PlaceOrder{
		Account: "eth1",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.BID,
		Rate:    2,
		Amount:  big.NewInt(2),
	})
	if err != nil {
		t.Fatalf("err placing order\n%v", err)
	}
	resO, err = obook.PlaceOrder(&PlaceOrder{
		Account: "eth2",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.BID,
		Rate:    2.5,
		Amount:  big.NewInt(3),
	})
	if err != nil {
		t.Fatalf("err placing order\n%v", err)
	}
	resO, err = obook.PlaceOrder(&PlaceOrder{
		Account: "eth3",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.BID,
		Rate:    3,
		Amount:  big.NewInt(2),
	})
	if err != nil {
		t.Fatalf("err placing order\n%v", err)
	}

	// an order should be filled to the correct amount
	resO, err = obook.PlaceOrder(&PlaceOrder{
		Account: "eos1",
		Symbol:  symbol.EOS_ETH,
		Type:    ordertype.ASK,
		Rate:    2.7,
		Amount:  big.NewInt(8),
	})
	if err != nil {
		t.Fatalf("err placing order\n%v", err)
	}
	if resO.ID == nil {
		t.Fatal("res id == nil")
	}

	if resO.Filled.String() != "5" {
		t.Fatalf("expected filled to be 5; received %s", resO.Filled.String())
	}

	var remaining string
	if err = db.QueryRow(`
		SELECT
			quantity

		FROM
			orderbook

		WHERE
			id = $1
	`, resO.ID).Scan(&remaining); err != nil {
		t.Fatalf("err querying remaining\n%v", err)
	}

	if remaining != "3" {
		t.Fatalf("expected remaining to be 3; received %s", remaining)
	}
}
