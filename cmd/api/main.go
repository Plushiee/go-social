package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/sikozonpc/social/internal/db"
	"github.com/sikozonpc/social/internal/env"
	"github.com/sikozonpc/social/internal/env/store"
)

func main() {
	godotenv.Load()
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgress://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenconns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_Idle_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenconns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic("Failed to connect to the database:", err)
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
