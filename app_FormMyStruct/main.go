package main

import "github.com/gin-gonic/gin"

type StructA struct {
	FileA string `form:"filed_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"filed_b"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func main()  {
	// 绑定表单数据至自定义结构体
	router := gin.Default()
	router.GET("/", GetDataB)
	router.Run("127.0.0.1:8080")
}
