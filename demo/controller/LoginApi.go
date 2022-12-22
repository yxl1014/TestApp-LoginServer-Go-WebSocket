package controller

import (
	"github.com/gin-gonic/gin"
)

func test_(c *gin.Context) {
	un := c.Query("username")
	data := c.Query("data")
	c.JSON(200, gin.H{
		"username": un,
		"data":     data,
	})
}
func test__(c *gin.Context) {
	un := c.Query("username")
	data := c.Query("data")
	c.ProtoBuf(200, gin.H{
		"username": un,
		"data":     data,
	})
}

func LoginController() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", test_)
	r.POST("/test", test__)
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run("0.0.0.0:11111")
}
