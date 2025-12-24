package main

import (
	"log"
	"smart-url/internal/env"
	"smart-url/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3080"),
	}

	store := store.NewStorage(nil)
	app := &application{
		config:  cfg,
		storage: store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
