package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/sikozonpc/social/internal/db"
	"github.com/sikozonpc/social/internal/env"
	"github.com/sikozonpc/social/internal/env/store"
)

const version = "0.0.1"

func main() {
	godotenv.Load()
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminPassword@localhost/social_db?sslmode=disable"),
			maxOpenconns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_Idle_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenconns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Connected to the database successfully")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
