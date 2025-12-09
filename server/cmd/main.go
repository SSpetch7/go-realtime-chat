package main

import (
	"log"
	"realtime_chat_server/db"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
}
