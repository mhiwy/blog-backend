package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	v0 := r.Group("/api")
	{
		v1 := v0.Group("/todo")
		{
			v1.Use(authorization)
			v1.POST("/add", addTodo)
			v1.POST("/update", updateTodo)
			v1.POST("/delete", deleteTodo)
			v1.POST("/list", getTodoList)
		}
		v2 := v0.Group("/user")
		{
			v2.POST("/in", signIn)
			v2.POST("/up", signUp)
			v2.POST("/out", signOut)
		}
		v3 := v0.Group("/post")
		{
			v3.Use(authorization)
			v3.POST("/add", addPost)
			v3.POST("/list", getPostList)
			v3.GET("/get", getPost)
		}
		v4 := v0.Group("/label")
		{
			v4.Use(authorization)
			v4.POST("/add", addLabel)
			v4.POST("/list", getLabelList)
		}
		v5 := v0.Group("/file")
		{
			v5.Use(authorization)
			v5.POST("/upload", uploadFile)
			v5.GET("/get", getFile)
			v5.GET("/list", getFileList)
			v5.GET("/delete", deleteFile)
		}
	}
	return r
}
