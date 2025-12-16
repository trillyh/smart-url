package main

import (
	"log"
	"smart-url/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3080"),
	}

	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
