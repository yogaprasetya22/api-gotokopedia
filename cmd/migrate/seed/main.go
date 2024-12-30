package main

import (
	"log"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost/gotokopedia?sslmode=disable")

	log.Println("Database address:", addr)

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}

