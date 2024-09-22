package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	url = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)
	os.Exit(m.Run())
}
