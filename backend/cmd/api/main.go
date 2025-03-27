package main

import (
	"log"
	"toodeloo/internal/env"
	"toodeloo/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	addr := env.GetString("ADDR", ":8080")
	cfg := config{
		addr: addr,
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
