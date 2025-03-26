package main

import (
	"log"
	"toodeloo/internal/env"

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

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
