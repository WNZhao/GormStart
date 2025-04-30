package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化gin的server对象
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8083 ") // listen and serve on 0.0.0.0:8080
}
