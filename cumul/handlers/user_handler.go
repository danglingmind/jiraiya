package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jiraiya/cumul/models"
	"github.com/jiraiya/cumul/utils"
)

// CheckUser : check if user exists
func CheckUser(c *gin.Context) {
	userid := c.Param("userid")
	userid = strings.Trim(userid, " ")
	userid = strings.ToLower(userid)
	if models.UserExists(userid) {
		c.JSON(200, gin.H{
			"message": "user exists",
		})
	} else {
		c.JSON(500, gin.H{
			"message": "user does not exists",
		})
	}
}

// NewUser : create the valid userid
func NewUser(c *gin.Context) {
	useridParam := c.Param("userid")
	passwordParam := c.Param("password") // TODO : decrypt
	useridParam = strings.ToLower(useridParam)

	if !utils.ValidateUserID(useridParam) {
		c.JSON(500, gin.H{
			"message": "userid must be less than 5 characters/invalid characters used",
			"err":     errors.New("UserID length exceeds"),
		})
	} else {
		// insert into table
		_, err := models.AddUser(useridParam, passwordParam)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "User Created",
				"error":   nil,
			})
		}
	}
	return
}

// Login the user
func Login(c *gin.Context) {
	userid := c.Param("userid")
	password := c.Param("password")
	if !utils.ValidateUserID(userid) {
		c.JSON(500, gin.H{
			"message": "userid must be less than 5 characters/invalid characters used",
			"err":     errors.New("Invalid userid"),
		})
	} else {
		loggedin := models.UserLogin(userid, password)
		if loggedin {
			c.JSON(200, gin.H{
				"message": "Login Successful",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "Login Failed",
			})
		}
	}
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
			c.JSON(500, gin.H{
				"message": "Please provide data",
			})
		} else {
			fmt.Println(urlsFromReq)
			added, err := models.StoreUrls(useridParam, urlsFromReq)
			if !added {
				fmt.Println(err.Error())
				c.JSON(500, gin.H{
					"err": err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"err":     nil,
					"message": "URLs added Successfully",
				})
			}
		}
	} else { // if user does not exists
		c.JSON(500, gin.H{
			"err": "Requested user does not exists",
		})
	}
}

// UserHome : User's Home
func UserHome(c *gin.Context) {

	useridParam := c.Param("userid")
	// check if the user exists
	if !models.UserExists(useridParam) {
		c.JSON(500, gin.H{
			"err":     "User Not Found",
			"message": "Requested user does not exists",
		})
	} else {
		// url struct to parse into html
		type urlstruct struct {
			Urlname string
			URL     string
		}
		urlmapout := make([]urlstruct, 0)
		urlMap, err := models.FetchUrls(useridParam)
		if err != nil {
			c.JSON(500, gin.H{
				"err": err.Error(),
			})
		} else {
			if len(urlMap) == 0 {
				// render home page with no url stored
				c.HTML(200, "no_url_home.tmpl", gin.H{
					"user": useridParam,
					"urls": urlmapout,
				})
			} else {
				for k, v := range urlMap {
					temp := urlstruct{Urlname: k, URL: v}
					urlmapout = append(urlmapout, temp)
				}
				c.HTML(200, "home1.tmpl", gin.H{
					"user": useridParam,
					"urls": urlmapout,
				})
			}
		}
	}
}

// UrlFetch return's all the urls in json
func UrlFetch(c *gin.Context) {
	useridParam := c.Param("userid")
	// check if the user exists
	if !models.UserExists(useridParam) {
		c.JSON(500, gin.H{
			"err":     "User Not Found",
			"message": "Requested user does not exists",
			"urls":    nil,
		})
	} else {
		// Fetch the urls for this user
		urlMap, err := models.FetchUrls(useridParam)
		if err != nil {
			c.JSON(500, gin.H{
				"err":     err.Error(),
				"message": nil,
				"urls":    nil,
			})
		} else {
			if len(urlMap) == 0 {
				c.JSON(200, gin.H{
					"err":     nil,
					"messege": "Zero URL available",
					"urls":    nil,
				})
			} else {
				// JSON output
				c.JSON(200, gin.H{
					"err":     nil,
					"message": nil,
					"urls":    urlMap,
				})
			}
		}
	}
}
