package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由器
	r := gin.Default()

	// 定义一个处理 GET 请求的路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Go; version=1")
	})

	// 在80端口上运行
	r.Run(":80")
}
