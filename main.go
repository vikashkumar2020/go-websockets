package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healhcheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Health check",
	})
}

func main() {
    router := gin.Default()
	router.GET("/",healhcheck)
    router.GET("/ws/:room", handleWebsocket)

    go broadcaster()

    if err := router.Run(":8080"); err != nil {
        log.Fatal("Unable to start server: ", err)
    }
}
