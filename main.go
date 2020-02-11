package main

import "net/http"

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", nil)
}
