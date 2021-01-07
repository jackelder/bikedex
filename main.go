package main

import (
	"fmt"
	"log"

	"github.com/jackelder/bikedex/data"
)

func main() {
	fmt.Println("running")

	db, err := data.Connect()

	if err != nil {
		log.Fatal(err)
	}

	ping := data.PingDatabase(db)

	if ping != nil {
		log.Fatal(ping)
	} else {
		fmt.Println("connected!")
	}

}
