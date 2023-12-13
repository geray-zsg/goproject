package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载HTML模板
	r.LoadHTMLGlob("./templates/*")
	// 定义GET方法路由
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "admin",
		})
	})
	r.Run(":8082")

}
