package main

import (
	"log"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}