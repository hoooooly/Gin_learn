package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default()

	// 提供unicode实体
	router.GET("/json", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"html":"<b>Hello world</b>",
		})
	})

	// 提供字面字符
	router.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(200, gin.H{
			"html":"<b>Hello world</b>",
		})
	})

	router.Run("127.0.0.1:8080")
}
