/*
 * @Author: holy
 * @Date: 2021-10-19 14:12:35
 * @LastEditTime: 2021-10-19 16:05:11
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Gin_learn\app\main.go
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()               // 创建默认的路由引擎
	r.GET("", func(c *gin.Context) { // 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
		c.JSON(200, gin.H{ //c.JSON：返回JSON格式的数据
			"message": "hello",
		})
	})

	//	 restful框架开发

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	// html模板渲染
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
			"name":  "holy",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "posts/index",
		})
	})

	// 自定义模板函数
	// r.SetFuncMap(template.FuncMap{
	// 	"safe": func(str string) template.HTML {
	// 		return template.HTML(str)
	// 	},
	// })
	// r.LoadHTMLFiles("./index.tmpl")

	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", "<a href='https://holychan.com'>Holy的博客</a>")
	// })

	r.Run() // 启动HTTP服务，默认在127.0.0.1:8080启动服务
}
