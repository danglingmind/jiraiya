package main

import (
	"github.com/gin-gonic/gin"
	cumulHandlers "github.com/jiraiya/cumul/handlers"
	"github.com/jiraiya/notabot"
)

// RoutesInit : Initialize all the routes
func RoutesInit() {
	// to test the availability
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// random arithmatics with ans
	router.GET("/notabot", notabot.RandomArithmatics)

	// cumul app
	router.GET("/cumul/:userid", cumulHandlers.User)
	router.GET("/cumul/new/:userid", cumulHandlers.NewUser)
}
