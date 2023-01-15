package controller

import (
	"fmt"
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
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))

		c.Abort()
		return
	}

	c.JSON(200, srv.GetFromWeb())
}
func (ctrl *rssController) CreateFeed(c *gin.Context) {
	var srv = service.CreateFeedService{}

	e := c.ShouldBind(&srv)
	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
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
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	c.JSON(200, srv.Get())
}
func (ctrl *rssController) DeleteFeed(c *gin.Context) {
	var srv = service.DeleteFeedService{}

	e := c.ShouldBind(&srv)
	if e != nil {
		fmt.Println("e", e)
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	c.JSON(200, srv.Delete())
}
func (ctrl *rssController) GetTodayItemList(c *gin.Context) {
	var srv = service.GetItemListService{}

	e := c.ShouldBindQuery(&srv)

	srv.PPagination = *goober.NewPagination().Querys(srv.PaginationQuery)

	if e != nil {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}
	c.JSON(200, srv.GetToday())
}

func (ctrl *rssController) UpdateFeed(c *gin.Context) {
	var srv = service.UpdateFeedItemsService{}
	href, h := c.GetQuery("href")

	if !h {
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, "缺少href"))
		c.Abort()
		return
	}

	c.JSON(200, srv.UpdateSinge(href))
}
func (ctrl *rssController) UpdateAllFeed(c *gin.Context) {
	var srv = service.UpdateFeedItemsService{}

	c.JSON(200, srv.UpdateAll())
}

var RssController = rssController{}
