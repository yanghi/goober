package main

import (
	"goblog/api"
	mysql "goblog/database/mysql"

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
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
