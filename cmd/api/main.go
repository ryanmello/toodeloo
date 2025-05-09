package main

import (
	"log"
	"toodeloo/internal/db"
	"toodeloo/internal/env"
	"toodeloo/internal/store"

	"github.com/joho/godotenv"
)

const version = "0.0.1"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("GOOSE_DBSTRING", "postgres://user:adminpassword@localhost/toodeloo?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_OPEN_CONNS", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
