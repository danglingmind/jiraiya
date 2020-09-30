package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiraiya/notabot"
)

var router *gin.Engine

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	fmt.Println("Hello, world !")
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/notabot", notabot.RandomArithmatics)

	router.Run(":8000")
}
