package fileTemplates

import "fmt"

type Main interface {
	MainContent(string) ([]string, []string)
}

type NetHttpMain struct{}

func (n *NetHttpMain) MainContent(modName string) ([]string, []string) {
	return []string{"main.go"}, []string{fmt.Sprintf(`package main

import (
	"log"
	"%s/server"
)

func main() {
	server := server.CreateServer()
	log.Fatal(server.ListenAndServe())
}
	`, modName)}
}
