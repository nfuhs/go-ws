package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewWebSocket(w, r)
		if err != nil {
			panic(err)
		}
		ws.On("message", func(e *Event) {
			log.Printf("[MESSAGE] %v", e.Data)
			ws.Out <- (&Event{
				Name: "response",
				Data: e.Data,
			}).Raw()
		})

	})
	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
