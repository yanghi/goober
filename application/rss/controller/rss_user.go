package controller

import (
	"goober/application/rss/service/user"
	"goober/auth"
	"goober/goober"

	"github.com/gin-gonic/gin"
)

type rssUserController struct {
}

func (ctrl *rssUserController) CreateFeed(c *gin.Context) {
	var srv = user.CreateFeedService{}

	e := c.ShouldBind(&srv)

	var u, _ = c.Get("user")

	srv.UserId = int(u.(*auth.JwtUserClaims).Uid)

	if e != nil || (srv.Url == "" && srv.FeedId == 0) {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))

		c.Abort()
		return
	}

	if srv.FeedId > 0 {
		c.JSON(200, srv.CreateWithFeedId())
	} else {
		c.JSON(200, srv.CreateWithUrl())
	}
}
func (ctrl *rssUserController) GetAllFeedList(c *gin.Context) {
	var srv = user.GetFeedListService{}

	var u, _ = c.Get("user")

	srv.UserId = int(u.(*auth.JwtUserClaims).Uid)

	c.JSON(200, srv.GetAll())
}
func (ctrl *rssUserController) MarkFeedItemReadState(c *gin.Context) {
	var srv = user.ModifyFeedService{}

	e := c.ShouldBind(&srv)

	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}
	var u, _ = c.Get("user")

	srv.UserId = int(u.(*auth.JwtUserClaims).Uid)

	c.JSON(200, srv.MarkReadState())
}

var RssUserController = rssUserController{}
