package fileTemplates

func handlerContent() (string, string) {
	return "handlers/handler.go", `package handlers

import (
	"html/template"
	"net/http"
)

var tmp *template.Template = template.Must(template.ParseFiles("templates/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	tmp.Execute(w, nil)
}
	`
}
