package main

import (
	"goblog/api"
	mysql "goblog/database/mysql"
	"goblog/router/middlewares"

	// serv "goblog/service"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// 基本用户信息
type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Avtar string `json:"avtar"`
}

type UserParams struct {
	Uid int `json:"uid"`
}

func main() {
	mysql.New()
	r := gin.Default()
	// r := gin.New()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/register", api.Register)
		v1.POST("/user/login", api.Login)
		v1.GET("/post/post", api.GetPost)
		v1.GET("/post/list", api.GetPostList)

		needAuth := v1.Group("/", middlewares.GetUserByToken(), middlewares.AuthRequired())

		needAuth.POST("post/post", api.CreatePost)
		needAuth.DELETE("post/post", api.DeletePostByAuthor)
		needAuth.PUT("post/post", api.ModifyPost)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}