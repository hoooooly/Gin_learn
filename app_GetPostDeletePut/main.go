package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	// GET请求
	router.GET("/path/:id", func(context *gin.Context) {
		// 获取路径参数参数id
		id := context.Param("id")
		// 获取查询参数pwd
		pwd := context.Query("pwd")
		// 设置默认的查询参数user
		user := context.DefaultQuery("user", "holy")
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})

	// POST请求
	router.POST("/path", func(context *gin.Context) {
		user := context.DefaultPostForm("user", "admin") // 默认是Admin用户
		pwd := context.PostForm("pwd")
		context.JSON(http.StatusOK, gin.H{
			"user":    user,
			"pwd":     pwd,
		})
	})

	// DELETE请求
	router.DELETE("/path/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"id":id,
		})
	})

	// PUT请求
	router.PUT("/path", func(context *gin.Context) {
		user := context.DefaultPostForm("user", "admin") // 默认是Admin用户
		pwd := context.PostForm("pwd")
		context.JSON(http.StatusOK, gin.H{
			"user":    user,
			"pwd":     pwd,
		})
	})

	router.Run("127.0.0.1:8080")
}
