package main

import (

	"github.com/gorilla/websocket"
)

type Room struct {
    Name    string
    Clients map[*websocket.Conn]Client
}

var rooms = make(map[string]Room)

func addClientToRoom(client Client) {
    room := rooms[client.room]
    if room.Clients == nil {
        room.Clients = make(map[*websocket.Conn]Client)
    }
    room.Clients[client.conn] = client
    rooms[client.room] = room
}

func removeClientFromRoom(client Client) {
    room := rooms[client.room]
    delete(room.Clients, client.conn)
    rooms[client.room] = room
}
