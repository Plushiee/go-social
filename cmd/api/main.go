package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/sikozonpc/social/internal/env"
	"github.com/sikozonpc/social/internal/env/store"
)

func main() {
	godotenv.Load()
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
