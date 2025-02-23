package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"com.sal/simple_bank/api"
	db "com.sal/simple_bank/db/sqlc"
	"com.sal/simple_bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {



	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("Cannot start db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start db:", err)
	}
	

}
