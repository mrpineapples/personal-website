package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool
var DBContext = context.Background()

func InitDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
	)
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
