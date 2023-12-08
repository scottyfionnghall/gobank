package main

import (
	"log"
)

func main() {
	connStr := "user=postgres dbname=gobank password=gobank sslmode=disable"

	store, err := NewPostgressStore(connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
		return
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
