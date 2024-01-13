package fileTemplates

import "fmt"

func mainContent(modName string) (string, string) {
	return "main.go", fmt.Sprintf(`package main

import (
	"log"
	"%sserver"
)

func main() {
	server := server.CreateServer()
	log.Fatal(server.ListenAndServe())
}
	`, modName)
}
