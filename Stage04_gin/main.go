package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main01() {
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

// 定义回调函数
// POST
func posting(c *gin.Context) {
	c.String(http.StatusOK, "POST")
}

// GET
func getting(c *gin.Context) {
	c.String(http.StatusOK, "GET")
}

// PUT
func putting(c *gin.Context) {
	c.String(http.StatusOK, "PUT")
}

// DELETE
func deleting(c *gin.Context) {
	c.String(http.StatusOK, "DELETE")
}

// PATCH
func patching(c *gin.Context) {
	c.String(http.StatusOK, "PATCH")
}

// HEAD
func heading(c *gin.Context) {
	c.String(http.StatusOK, "PATCH")
}

// OPTIONS
func optioning(c *gin.Context) {
	c.String(http.StatusOK, "optioning")
}

func main02() {
	// 1.创建路由
	router := gin.Default()

	// 2.绑定路由规则
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDel", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", heading)
	router.OPTIONS("/someOption", optioning)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func main03() {
	router := gin.Default()
	// 此规则能够匹配/user/freshen这种格式，但不能匹配/user/或/user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此规则能够匹配/user/freshen/格式也能匹配/user/freshen/send这种格式
	// 如果没有其他路由器匹配/user/freshen,它将重定向到/user/freshen/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func main04() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		// 指定默认值
		name := c.DefaultQuery("name", "Gopher")
		c.String(http.StatusOK, "Hello %s", name)
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
func main05() {
	router := gin.Default()
	// 解析/路径给用户返回默认首页
	// 加载html目录下所有静态页面
	router.LoadHTMLGlob("html/*")
	// GET请求访问/路径返回index.html
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "文心雕龙"})
	})

	router.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, "type:%s ,name:%s,password: %s", types, username, password)
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func main() {
	router := gin.Default()
	// 解析/路径给用户返回默认首页
	// 加载HTML文件
	router.LoadHTMLGlob("html/*")
	// GET请求访问/路径返回upload.html
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"说明": "返回给用户upload.html页面"})
	})
}
