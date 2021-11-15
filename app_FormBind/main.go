package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostParams struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  bool   `json:"sex" form:"sex"`
}

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	// 绑定JSON
	router.POST("/post", func(context *gin.Context) {
		var p PostParams                  // 实例化
		err := context.ShouldBindJSON(&p) // 将地址传入ShouldBindJSOn，接收JSON格式的参数
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg": err,
				"data": gin.H{},
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "success",
				"data": p,
			}) // 没有错误直接抛出p
		}
	})
	// 绑定URI
	// 127.0.0.1:8080/person/holy/987fbc97-4bed-5078-9f07-9141ba07c9f3
	router.GET("/person/:name/:id", func(context *gin.Context) {
		var person Person
		if err:= context.ShouldBindUri(&person); err!= nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(), // 返回错误
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	// 绑定表单
	router.POST("/form", func(context *gin.Context) {

	})

	router.Run("127.0.0.1:8080")
}
