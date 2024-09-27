package db

import (
	"context"
	"log"
	"os"
	"testing"

	"simple_bank/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)
	os.Exit(m.Run())
}
