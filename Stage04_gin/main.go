package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	router := gin.Default()
	// 2.绑定路由规则，执行函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	//3.监听端口，默认在8080
	err := router.Run()
	if err != nil {
		return
	}
}
