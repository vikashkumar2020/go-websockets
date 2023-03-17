package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebsocket(c *gin.Context) {
    room := c.Param("room")
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Fatal("Unable to upgrade connection: ", err)
    }

    client := Client{conn, "", room}
    entering <- client

    for {
        var message Message
        if err := conn.ReadJSON(&message); err != nil {
            fmt.Print("err",err)
            leaving <- client
            break
        }

        messages <- Message{client.name, message.Content}
    }
}
