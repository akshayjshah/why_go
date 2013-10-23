package main

import (
	"fmt"
	"github.com/hearsaycorp/gochat"
)

// START OMIT
func main() {
	eng := gochat.NewListener("Engineering")
	go eng.Listen()

	infra := gochat.NewListener("Infrastructure")
	go infra.Listen()

	for {
		select {
		case msg := <-eng.Messages:
			fmt.Println("Engineering: ", msg.Message)
		case msg := <-infra.Messages:
			fmt.Println("Infrastructure: ", msg.Message)
		}
	}
}

// END OMIT
