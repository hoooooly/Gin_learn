<!--
 * @Author: holy
 * @Date: 2021-10-19 14:06:57
 * @LastEditTime: 2021-10-19 16:05:26
 * @LastEditors: holy
 * @Description: 
 * @FilePath: \Gin_learn\README.md
-->
# Gin_learn

gin框架学习笔记和代码

## 安装

```bash
$ go get -u github.com/gin-gonic/gin
```

### 快速开始

```bash
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()               // 创建默认的路由引擎
	r.GET("", func(c *gin.Context) { // 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
		c.JSON(200, gin.H{ //c.JSON：返回JSON格式的数据
			"message": "hello",
		})
	})
	r.Run() // 启动HTTP服务，默认在127.0.0.1:8080启动服务
}
```

访问127.0.0.1:8080即可访问服务。

## RESTful API

REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

推荐阅读阮一峰老师的[理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)。

- GET用来获取资源
- POST用来新建资源
- PUT用来更新资源
- DELETE用来删除资源。

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

Gin框架支持开发RESTful API的开发。

```bash
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
```

开发RESTful API的时候我们通常使用Postman来作为客户端的测试工具。

## Gin渲染

### HTML渲染

首先定义一个存放模板文件的templates文件夹，然后在其内部按照业务分别定义一个posts文件夹和一个users文件夹。 posts/index.html文件的内容如下：

```bash
{{define "posts/index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。

```bash
	// html模板渲染
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
```

<!-- ### 自定义模板函数

定义一个不转义相应内容的safe模板函数如下：

```bash
func main() {
	router := gin.Default()
	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://holychan.com'>Holy的博客</a>")
	})

	router.Run(":8080")
}
```

在index.tmpl中使用定义好的safe模板函数：

```bash
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>修改模板引擎的标识符</title>
</head>
<body>
<div>{{ . | safe }}</div>
</body>
</html>
```
 -->



