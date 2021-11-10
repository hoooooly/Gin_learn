package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/uploadsfile", func(context *gin.Context) {
		context.HTML(http.StatusOK, "uploads.html", gin.H{

		})
	})
	
	// 设置最低内存限制，默认是32MiB
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// 单个文件
		file, err := context.FormFile("file")
		if err != nil {
			context.String(500, "上传文件出错")
		}
		log.Println(file.Filename)

		// 上传文件到特定地址

		context.SaveUploadedFile(file, "./file/" + file.Filename)
		context.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
	})
	router.Run()
}
