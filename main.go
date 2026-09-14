package main

import (
	"log"

	"github.com/Dmitry-Dyagilev/final-project/pkg/db"
	"github.com/Dmitry-Dyagilev/final-project/pkg/server"
)

func main() {

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	port := ":7540"
	if err := server.Run(port); err != nil {
		log.Fatal(err)
	}
}
