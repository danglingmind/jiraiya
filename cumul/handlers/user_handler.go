package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiraiya/cumul/models"
)

// NewUser : create the valid userid
func NewUser(c *gin.Context) {
	useridParam := c.Param("userid")
	// TODO : validate userid

	// insert into table
	_, err := models.AddUser(useridParam)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Created",
		"error":   nil,
	})
}

// UserURLStore : stores all the urls
func UserURLStore(c *gin.Context) {
	useridParam := c.Param("userid")

	if models.UserExists(useridParam) {
		// iterate over all the urls and store them into db
		// this should be moved out of here and added into Async Queue
		urlsFromReq := make(map[string]string)
		c.BindJSON(&urlsFromReq)

		// if no url given
		if len(urlsFromReq) == 0 {
			c.JSON(200, gin.H{
				"message": "Please provide data",
			})
			return
		}

		added, err := models.StoreUrls(useridParam, urlsFromReq)
		if !added {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"err":     nil,
			"message": "URLs added Successfully",
		})
	} else { // if user does not exists
		c.JSON(500, gin.H{
			"err": "Requested user does not exists",
		})
	}
}

// UserURLFetch : fetches all the urls for this user
func UserURLFetch(c *gin.Context) {
	useridParam := c.Param("userid")
	// check if the user exists
	if !models.UserExists(useridParam) {
		c.JSON(500, gin.H{
			"err":     "User Not Found",
			"message": "Requested user does not exists",
		})
		return
	}
	// Fetch the urls for this user
	urlMap, err := models.FetchUrls(useridParam)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	if len(urlMap) == 0 {
		c.JSON(200, gin.H{
			"err":     nil,
			"messege": "Zero URL available",
		})
		return
	}
	// JSON output
	// c.JSON(200, gin.H{
	// 	"urls": urlMap,
	// })

	// // render html
	type urlstruct struct {
		Urlname string
		URL     string
	}
	urlmapout := make([]urlstruct, 0)
	for k, v := range urlMap {
		temp := urlstruct{Urlname: k, URL: v}
		urlmapout = append(urlmapout, temp)
	}
	fmt.Println(urlmapout)
	c.HTML(200, "home1.tmpl", gin.H{
		"user": useridParam,
		"urls": urlmapout,
	})
}
