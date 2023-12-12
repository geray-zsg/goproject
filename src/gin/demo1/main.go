package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 定义路由的GET方法及相应处理函数
	r.GET("/hello", func(ctx *gin.Context) {
		// 将发送的信息封装成JSON发送给浏览器
		ctx.JSON(http.StatusOK, gin.H{
			// 这里定义我们的数据
			"message": "快速入门：https://www.w3cschool.cn/golang_gin/golang_gin-3vd23lry.html",
		})
	})
	// 启动服务
	r.Run(":8088")
}
