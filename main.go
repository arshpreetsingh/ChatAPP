package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type clinet struct {

	// web socker for communication
	socket *websocket.Conn

	// send is channel on which message will be sent

	send chan []byte

	// room for client to stay in, Will be of type struct
	room *room
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
