package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiraiya/cumul/models"
)

// UserFetch : fetches all the urls for this user
func UserFetch(c *gin.Context) {
	var user models.User
	useridParam := c.Param("userid")

	// get the username
	err := models.DB.QueryRow("select userid from user where userid=?", useridParam).Scan(&user)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "User not found",
		})
	} else {
		// Fetch the urls for this user
		results, err := models.DB.Query("Select url from urls where userid = ?", useridParam)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "Could not find urls for this user",
			})
		}
		defer results.Close()

		urls := make([]models.Url, 0)
		for results.Next() {
			var url models.Url
			if err := results.Scan(&url); err != nil {
				c.JSON(500, gin.H{
					"error":   err.Error(),
					"message": "Could not fetch the urls for this user",
				})
			}
			urls = append(urls, url)
		}
		// TODO : prepare a html page with these URLS
		// TODO : prepare a JSON for these urls

	}
}

// NewUser : create the valid userid
func NewUser(c *gin.Context) {
	useridParam := c.Param("userid")
	// TODO : validate userid
	// insert into table
	stmt, err := models.DB.Prepare("INSERT INTO user (userid, paid) values (?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(useridParam)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	err = models.DB.Commit()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Created",
		"error":   nil,
	})
}
