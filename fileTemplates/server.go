package fileTemplates

import "fmt"

func serverContent(modName string) (string, string) {
	return "server/server.go", fmt.Sprintf(`package server

import (
	"net/http"
	"time"

	"%s/handlers"
	"github.com/gorilla/mux"
)

func CreateServer() *http.Server {
	mux := mux.NewRouter()
	// 用filerServer去存取static folder
	fs := http.FileServer(http.Dir("./static"))
	// mux去handle關於static folder的調用，StripPrefix把/static/的前綴拿掉，不然會變成./static/static/xxx.css
	mux.Handle("/static/", http.StripPrefix("/static/", fs)) // 注意寫法: /static/
	mux.HandleFunc("/", handlers.Home)
	server := http.Server{
		Addr:         ":5000",
		Handler:      mux,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}

	return &server
}
	`, modName)
}
