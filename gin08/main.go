package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("example", "12345")
		context.Next()
		latency := time.Since(t)
		fmt.Printf("%s - %s - %s - %s\n", context.ClientIP(), context.Request.Method, context.Request.URL, latency)
		fmt.Printf("%d\n", context.Writer.Status())
	}
}

func main() {
	//router := gin.New()
	// 使用logger中间件
	//router.Use(gin.Logger())
	// 使用recovery中间件
	//router.Use(gin.Recovery())
	router := gin.Default()
	// 以上是使用default的方式是一样的
	authrized := router.Group("/auth")
	authrized.Use(MyLogger())
	authrized.GET("/ping", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		context.JSON(http.StatusOK, gin.H{
			"message": example,
		})
	})
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8083")
}
