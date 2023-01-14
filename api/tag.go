package api

import (
	"fmt"
	"goober/goober"
	tag_service "goober/service/tag"

	"github.com/gin-gonic/gin"
)

func TagCreate(c *gin.Context) {
	var service tag_service.CreateTagService

	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	c.JSON(200, service.Create())
}
func TagDelete(c *gin.Context) {
	var service tag_service.DeleteTagService

	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	c.JSON(200, service.Del())
}
func TagModify(c *gin.Context) {
	var service tag_service.ModifyTagService

	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, goober.FailedResult(goober.ErrParamsInvlid, ""))
		c.Abort()
		return
	}

	c.JSON(200, service.Modify())
}

func TagGetList(c *gin.Context) {
	var service tag_service.GetTagListService

	c.JSON(200, service.Get())
}
