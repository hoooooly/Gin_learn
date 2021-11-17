package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.POST("/testUpload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.SaveUploadedFile(file, )
		context.JSON(200, gin.H{
			"msg": file,
		})
	})

	r.Run(":8080")
}
