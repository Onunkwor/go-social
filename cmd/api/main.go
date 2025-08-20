package main

import (
	"context"
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
	// dynamic database selection
	envName := env.GetString(os.Getenv("APP_ENV"), "dev")
	var dbName string
	switch envName {
	case "test":
		dbName = "social_test"
	case "prod":
		dbName = "social"
	default:
		dbName = "social_dev"
	}
	cfg := config{
		addr: env.GetString(addr, ":4000"),
		db: dbConfig{
			dsn: env.GetString(os.Getenv("MONGODB_URI"), ""),
		},
	}
	client, db, err := db.Connect(cfg.db.dsn, dbName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	store := store.NewStorage(db)
	defer client.Disconnect(context.Background())

	// log.Printf("Connected to database at %s", cfg.db.dsn)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
