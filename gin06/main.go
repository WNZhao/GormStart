package main

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"GormStart/gin06/proto"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "gin"
		msg.Message = "hello"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	//protobuf
	router.GET("/moreProtoBuf", func(c *gin.Context) {
		c.ProtoBuf(http.StatusOK, &proto.Teacher{
			Name:   "gin",
			Course: []string{"go", "python"},
		})
	})

	router.Run(":8083")
}
