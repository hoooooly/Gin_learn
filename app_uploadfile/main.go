package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main()  {
	r := gin.Default()
	r.POST("/testUpload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		name := context.PostForm("name")
		// context.SaveUploadedFile(file, "./files/" + file)
		in, _ := file.Open()
		defer file.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		context.JSON(200, gin.H{
			"msg": file,
			"name": name,
		})
	})

	r.Run(":8080")
}
