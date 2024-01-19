package fileTemplates

type Handler interface {
	HandlerContent() ([]string, []string)
}

type NetHttpHandler struct{}

func (n *NetHttpHandler) HandlerContent() ([]string, []string) {
	return []string{"handlers/handler.go"}, []string{`package handlers

import (
	"html/template"
	"net/http"
)

var tmp *template.Template = template.Must(template.ParseFiles("templates/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	tmp.Execute(w, nil)
}
	`}
}
