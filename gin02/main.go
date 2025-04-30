package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	// restful 的开发中
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// 默认启动的是 8080 端口，也可以通过 router.Run(":端口号") 自定义
	router.Run()

}

func options(context *gin.Context) {

}

func head(context *gin.Context) {

}

func deleting(context *gin.Context) {

}

func patching(context *gin.Context) {

}

func putting(context *gin.Context) {

}

func posting(context *gin.Context) {

}

func getting(context *gin.Context) {

}
