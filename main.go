package main

import (
	"fmt"
	"goober/api"
	"goober/application/rss/controller"
	rss_service "goober/application/rss/service"
	"goober/config"
	"goober/database/mysql"
	"goober/goober"
	"goober/router/middlewares"
	"os"

	// serv "goober/service"
	// "net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
}

func main() {

	var r *gin.Engine
	var mode = os.Getenv("GB_MODE")
	fmt.Println("use env", os.Getenv("GB_MODE"))
	// load configuration
	config.LoadConfig()

	mysql.New(config.AppConf.MySql)

	if mode == "prod" || config.AppConf.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	}
	goober.Logger().Info("[goober] app ready")

	var feedJob = rss_service.FeedUpdateJobService{}
	feedJob.Start()

	var usePprof = config.AppConf.Debug.Pprof

	if usePprof {
		pprof.Register(r)
	}

	r.Use(middlewares.Cors())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/register", api.Register)
		v1.POST("/user/login", api.Login)
		v1.GET("/post/post", middlewares.TryGetUserByToken(), api.GetPost)
		v1.GET("/post/list", api.GetPostList)
		v1.PUT("/post/action/view", api.PostActionView)

		v1.GET("/rss/feed/web", controller.RssController.GetFeed)
		v1.GET("/rss/feed/all", controller.RssController.GetAllFeed)
		v1.POST("/rss/feed", controller.RssController.CreateFeed)
		v1.DELETE("/rss/feed", controller.RssController.DeleteFeed)
		v1.POST("/rss/feed/update", controller.RssController.UpdateFeed)
		v1.PUT("/rss/feed", controller.RssController.ModifyFeed)
		v1.POST("/rss/feed/updateAll", controller.RssController.UpdateAllFeed)
		v1.GET("/rss/item/list", controller.RssController.GetItemList)
		v1.GET("/rss/item/today", controller.RssController.GetTodayItemList)

		needAuth := v1.Group("/", middlewares.GetUserByToken(), middlewares.AuthRequired())
		needAuth.POST("rss/user/feed", controller.RssUserController.CreateFeed)
		needAuth.GET("rss/user/feed/all", controller.RssUserController.GetAllFeedList)
		v1.GET("rss/user/item/list", controller.RssController.GetItemList)

		needAuth.GET("user/post/list", api.GetUserPostList)

		needAuth.GET("user/info/base", api.GetUserBaseInfo)
		needAuth.PUT("user/info/info", api.ModifyUserInfo)

		needAuth.GET("tag/list", api.TagGetList)
		needAuth.POST("tag/tag", api.TagCreate)
		needAuth.PUT("tag/tag", api.TagModify)
		needAuth.DELETE("tag/tag", api.TagDelete)

		needAuth.POST("post/post", api.CreatePost)
		needAuth.DELETE("post/post", api.DeletePostByAuthor)
		needAuth.PUT("post/post", api.ModifyPost)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
