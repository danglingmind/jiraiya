package routes

import (
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
func Init(router *gin.Engine) {
	// to test the availability
	router.GET("/ping", ping)

	// random arithmatics with ans
	router.GET("/notabot", notabot.RandomArithmatics)

	// cumul app
	cumul := router.Group("/cumul")
	{
		cumul.GET("/:userid", cumulHandlers.UserURLFetch)  // this will fetch all the URLs
		cumul.POST("/:userid", cumulHandlers.UserURLStore) // this will store the URLs
		cumul.GET("/:userid/new", cumulHandlers.NewUser)
	}

}
