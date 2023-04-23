package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "michaelsite_dev"
)

var Pool *pgxpool.Pool
var DBContext = context.Background()

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}
	Pool = pool
	fmt.Println("Models successfully connected to postgres!")
}
