package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			log.Println("Error creating websocket connection: %v", err)
			return
		}
		ws.On("message", func(e *Event) {
			log.Printf("Message received: %s", e.Data.(string))
			ws.Out <- (&Event{
				Name: "response",
				Data: strings.ToUpper(e.Data.(string)),
			}).Raw()
		})
	})

	log.Println("WebServer listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
