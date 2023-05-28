package main

import (
	"log"
	"test/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
