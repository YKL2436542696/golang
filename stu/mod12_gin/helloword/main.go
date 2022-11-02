package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 1. 创建路由
	r := gin.Default()
	// 2. 绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/user", func(c *gin.Context) {
		// 默认值
		name := c.DefaultQuery("name", "小月")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// 3. 监听端口，默认在8080
	r.Run(":8000")

}
