package main

import (
	"log"

	"github.com/Cody-Kao/default/server"
)

func main() {
	server := server.CreateServer()
	log.Fatal(server.ListenAndServe())
}
