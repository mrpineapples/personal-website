package database

import (
	"context"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool
var DBContext = context.Background()

func InitDB() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}
	Pool = pool
	log.Println("Database connection successful!")
}

func RunMigrations() {
	m, err := migrate.New(
		os.Getenv("MIGRATION_URL"),
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		log.Fatal("Error creating migrate instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error running up migrate:", err)
	}

	log.Println("Database migration successful")
}
