package controller

import (
	"goober/application/rss/service"
	"goober/goober"

	"github.com/gin-gonic/gin"
)

type rssController struct {
}

func (ctrl *rssController) GetFeed(c *gin.Context) {
	var srv = service.GetFeedService{}

	e := c.ShouldBindQuery(&srv)

	if e != nil || srv.Href == "" {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid))

		c.Abort()
		return
	}

	c.JSON(200, srv.GetFromWeb())
}
func (ctrl *rssController) CreateFeed(c *gin.Context) {
	var srv = service.CreateFeedService{}

	e := c.ShouldBind(&srv)
	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, srv.Create())
}
func (ctrl *rssController) GetAllFeed(c *gin.Context) {
	var srv = service.GetFeedListService{}

	c.JSON(200, srv.GetAll())
}

type F struct {
	goober.PaginationQuery
	FeedId int `json:"feedId" form:"feedId"`
}

func (ctrl *rssController) GetItemList(c *gin.Context) {
	var srv = service.GetItemListService{}

	e := c.ShouldBindQuery(&srv)

	srv.PPagination = *goober.NewPagination().Querys(srv.PaginationQuery)

	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, srv.Get())
}

var RssController = rssController{}
