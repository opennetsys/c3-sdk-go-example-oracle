// +build integration

package pg

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	// 1. connect to the db
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	TEST_POSTGRES_URL := os.Getenv("TEST_POSTGRES_URL")
	if TEST_POSTGRES_URL == "" {
		t.Fatalf("TEST_POSTGRES_URL env var is required")
	}

	pgstore, err := New(&Options{
		PostgresURL: TEST_POSTGRES_URL,
	})
	if err != nil {
		t.Fatalf("err new\n%v", err)
	}

	if err = pgstore.db.Ping(); err != nil {
		t.Fatalf("err pinging\n%v", err)
	}
}
