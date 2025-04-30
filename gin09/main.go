package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 创建一个 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	// 启动服务放到协程中
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// 创建退出通道监听系统信号
	quit := make(chan os.Signal, 1)
	// 监听中断信号或终止信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // 等待信号

	fmt.Println("正在优雅退出服务器...")

	// 创建上下文设置最大超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("服务器强制关闭:", err)
	} else {
		fmt.Println("服务器优雅退出")
	}
}
