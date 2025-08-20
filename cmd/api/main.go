package main

import (
	"github/onunkwor/social/internal/db"
	"github/onunkwor/social/internal/env"
	"github/onunkwor/social/internal/store"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("API_ADDR")
	cfg := config{
		addr: env.GetString(addr, ":4000"),
		db: dbConfig{
			dsn:          env.GetString(os.Getenv("DB_DSN"), "postgres://user:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt(os.Getenv("DB_MAX_OPEN_CONNS"), 25),
			maxIdleConns: env.GetInt(os.Getenv("DB_MAX_IDLE_CONNS"), 25),
			maxIdleTime:  env.GetString(os.Getenv("DB_MAX_IDLE_TIME"), "15m"),
		},
	}
	db, err := db.New(
		cfg.db.dsn,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	store := store.NewStorage(db)
	defer db.Close()

	// log.Printf("Connected to database at %s", cfg.db.dsn)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
