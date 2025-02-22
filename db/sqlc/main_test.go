package db

import (
	"context"
	"log"
	"os"
	"testing"

	"com.sal/simple_bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot Load config")
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
