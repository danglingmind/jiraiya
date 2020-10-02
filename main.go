package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiraiya/cumul/models"
	"github.com/jiraiya/cumul/routes"
)

var (
	router *gin.Engine
)

func main() {
	// initialize database
	models.InitDB()
	// setup router
	router := gin.Default()
	// initialize routes for endpoints
	routes.Init(router)
	router.Run(":8000")

}
