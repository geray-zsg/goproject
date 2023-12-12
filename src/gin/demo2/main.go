package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 定义路由GET方法：查询学生信息
	r.GET("/student", func(ctx *gin.Context) {
		// 封装信息成JSON发送给浏览器
		ctx.JSON(http.StatusOK, gin.H{
			"message": "查询学生信息成功",
		})
	})

	// 定义路由POST方法：创建学生信息
	r.POST("/createStudnet", func(ctx *gin.Context) {
		// 封装JSON信息发送给浏览器
		ctx.JSON(http.StatusOK, gin.H{
			"message": "创建学生信息成功",
		})
	})

	// 定义路由PUT方法：更新学生信息
	r.PUT("/upgradeStudent", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "更新学生信息成功",
		})
	})

	// 定义路由DELETE方法：删除学生信息
	r.DELETE("/deleteStudent", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "删除学生成功",
		})
	})

	r.Run(":8081")

}
