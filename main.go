package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/dron-ci-test/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong YHO-HO",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
