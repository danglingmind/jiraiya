package notabot

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RandomArithmatics returns a json with random question with ans
func RandomArithmatics(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(20)
	b := rand.Intn(20)

	randomOp := rand.Intn(4)

	// get ans
	var ans interface{}
	var que string
	switch randomOp {
	case 0:
		ans = a + b
		que = strconv.Itoa(a) + " + " + strconv.Itoa(b)
	case 1:
		ans = a - b
		que = strconv.Itoa(a) + " - " + strconv.Itoa(b)
	case 2:
		ans = a * b
		que = strconv.Itoa(a) + " * " + strconv.Itoa(b)
	case 3:
		// avoid devide by zero
		if b == 0 {
			b = 1
		}
		ans = a / b
		que = strconv.Itoa(a) + " / " + strconv.Itoa(b)
	}
	c.JSON(http.StatusOK, gin.H{
		"question": que,
		"answer":   ans,
	})
}
