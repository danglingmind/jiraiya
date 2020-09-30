package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getNewUser() string {
	return "ricky"
}

func startGame(c *gin.Context) {
	// if board exists then join this user

	// else create a new board
	// return a unique userid
}

func makeMove(c *gin.Context) {
	// if match is still alive
	//  and if move is valid then make move
	// if match ended then return who this user WON/LOST
}

func moveList(c *gin.Context) {
	// fetch all the move list
	// create json response and return it
}

func main() {
	fmt.Println("Running ...")

	router = gin.Default()
	// test the server
	router.GET("/ping", pong)
	// start new game
	router.POST("/game/start", startGame)
	// make a move
	router.POST("/game/move/:userid/:col", makeMove)
	// get the list of all the moves by both the users
	router.GET("/game/movelist", moveList)
}
