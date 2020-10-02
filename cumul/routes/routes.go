package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	cumulHandlers "github.com/jiraiya/cumul/handlers"
	"github.com/jiraiya/notabot"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Init : Initialize all the routes
func Init(router *gin.Engine, db *sql.DB) {
	// to test the availability
	router.GET("/ping", ping)

	// random arithmatics with ans
	router.GET("/notabot", notabot.RandomArithmatics)

	// cumul app
	router.GET("/cumul/:userid", cumulHandlers.UserFetch)
	router.GET("/cumul/:userid/new", cumulHandlers.NewUser)
}
