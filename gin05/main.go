package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 匹配的 URL 格式：/welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 等价于 c.Request.URL.Query().Get("lastname")

		c.JSON(http.StatusOK, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})

	router.POST("/welcome", func(c *gin.Context) {
		job := c.DefaultPostForm("job", "Guest")
		salary := c.PostForm("salary") // 等价于 c.Request.URL.Query().Get("lastname")

		//c.String(http.StatusOK, "Hello %s %s", job, salary)
		c.JSON(http.StatusOK, gin.H{
			"job":    job,
			"salary": salary,
		})

	})

	router.Run(":8083")
}
