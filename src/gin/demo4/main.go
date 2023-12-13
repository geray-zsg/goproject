// Gin框架可以获取参数

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// queryGet()
	formPost()
}

// 通过Path参数获取
func pathGet() {
	r := gin.Default()

	r.GET("/user/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
		})

	})
}

// 当前端请求的数据通过form表单提交时，例如向​/user/reset​发送了一个POST请求，获取请求数据方法如下
func formPost() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html") //加载页面

	// 定义GET请求路由
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	// 定义POST请求路由
	r.POST("/", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		fmt.Printf("username: %s; password: %s", username, password)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":8081")
}

// 获取Query参数：通过Query来获取URL的?后面所携带的参数，例如：/name=admin&pws=amdin
func queryGet() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {

		// 获取URL的参数
		name := ctx.Query("name")
		pwd := ctx.Query("pwd")
		fmt.Printf("name: %s; pwd: %s", name, pwd)

		// 将发送到浏览器的数据封装成JSON
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
	// 默认端口是8080
	r.Run(":8080")
}
