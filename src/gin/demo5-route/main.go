package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 共同前缀的URL路由划分为一个路由组
	user := r.Group("/user")

	// http://127.0.0.1:8080/user/index
	user.GET("/index", func(ctx *gin.Context) {
		// 封装JSON给浏览器
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET请求到index路径",
		})
	})
	// http://127.0.0.1:8080/user/login
	user.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "POST请求到login",
		})
	})

	// 支持嵌套路由: http://127.0.0.1:8080/user/pwd/pwd
	pwd := user.Group("/pwd")
	pwd.GET("/pwd", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "嵌套路由：http://127.0.0.1:8080/user/pwd/pwd",
		})
	})

	r.Run(":8080")
}
