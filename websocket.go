package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) vool {
		return true
	},
}

type Websocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

func NewWebsocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("An error occured while upgrading the connection: %v", err)
		return nil, err
	}
}
