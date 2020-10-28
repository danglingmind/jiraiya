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

	// load all the htmls
	router.LoadHTMLGlob("cumul/templates/*")
	// load static
	router.Static("assets/css", "cumul/assets/css")
	router.Static("assets/js", "cumul/assets/scripts")

	// cumul app
	cumul := router.Group("/cumul")
	{
		cumul.GET("/:userid", cumulHandlers.UserHome)              // this is the Home Screen
		cumul.GET("/:userid/urls", cumulHandlers.UrlFetch)         // returns all the stored URLs
		cumul.POST("/:userid", cumulHandlers.UserURLStore)         // this will store the URLs
		cumul.GET("/:userid/new/:password", cumulHandlers.NewUser) // add new user
		cumul.GET("/:userid/check", cumulHandlers.CheckUser)       // check if user exists
		cumul.GET("/:userid/login/:password", cumulHandlers.Login) // check if user exists
	}
	cumulrender := router.Group("render/cumul")
	{
		cumulrender.GET("/user/register", cumulHandlers.RegisterRender)
		cumulrender.GET("/user/inputurl/:userid", cumulHandlers.InputURL)
	}

}
