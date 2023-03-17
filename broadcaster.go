package main

import (
    "log"
)

var (
    messages = make(chan Message)
    entering = make(chan Client)
    leaving  = make(chan Client)
)

func broadcaster() {
    for {
        select {
        case message := <-messages:
            room := rooms[message.Sender]
            for _, client := range room.Clients {
                if client.name != message.Sender {
                    if err := client.conn.WriteJSON(message); err != nil {
                        log.Println("Unable to send message: ", err)
                    }
                }
            }

        case client := <-entering:
            addClientToRoom(client)
            log.Println("Client entered:", client.conn.RemoteAddr())

        case client := <-leaving:
            removeClientFromRoom(client)
            log.Println("Client left:", client.conn.RemoteAddr())
        }
    }
}
