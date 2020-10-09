package handlers

import "github.com/gin-gonic/gin"

func RegisterRender(c *gin.Context) {
	c.HTML(200, "register.tmpl", gin.H{})
}

func InputURL(c *gin.Context) {
	userid := c.Param("userid")

	c.HTML(200, "addurl.tmpl", gin.H{
		"user": userid,
	})
}
