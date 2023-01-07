package api

import (
	"fmt"
	"goblog/auth"
	gerr "goblog/error"
	"goblog/model"
	postmodel "goblog/model/post"
	"goblog/rep"
	"goblog/service/post"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var service post.CreatePostService

	var u, _ = c.Get("user")

	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)

	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, service.Run())
}

func DeletePostByAuthor(c *gin.Context) {
	var service post.DeletePostService

	var u, _ = c.Get("user")

	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)
	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	var id = c.Query("id")
	pid, e := strconv.Atoi(id)

	if e != nil {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}
	service.Id = pid

	c.JSON(200, service.DeleteByAuthor())
}

func GetPost(c *gin.Context) {
	var service post.GetPostService

	sid, hasId := c.GetQuery("id")

	if !hasId {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}
	id, er := strconv.Atoi(sid)

	if er != nil {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	service.Id = id

	u, _ := c.Get("user")

	if u != nil {
		service.UserId = int(u.(*auth.JwtUserClaims).Uid)
		c.JSON(200, service.GetByUser())
	} else {
		c.JSON(200, service.Get())
	}

}

func GetPostList(c *gin.Context) {
	var service post.GetPostListService

	pq := model.PaginationQuery{}

	c.ShouldBindQuery(&pq)
	pg := model.Pagination{}
	service.PaginationParams = *pg.Query(pq)

	c.JSON(200, service.Get())
}

func GetUserPostList(c *gin.Context) {
	var service post.GetPostListService

	var u, _ = c.Get("user")
	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)

	var statuQuery = c.Query("statu")

	if statuQuery != "" {
		statu, e := strconv.Atoi(statuQuery)
		if e == nil {
			service.Statu = postmodel.IToPostStatu(statu)
		} else {
			service.Statu = postmodel.PostStatuNotExsit
		}
	} else {
		service.Statu = postmodel.PostStatuNotExsit
	}

	pq := model.PaginationQuery{}

	c.ShouldBindQuery(&pq)
	pg := model.Pagination{}
	service.PaginationParams = *pg.Query(pq)

	c.JSON(200, service.GetByAuthor())
}

func ModifyPost(c *gin.Context) {
	var service post.ModifyPostService
	service.Statu = postmodel.PostStatuNotExsit

	var u, _ = c.Get("user")

	service.AuthorId = int(u.(*auth.JwtUserClaims).Uid)
	e := c.ShouldBind(&service)
	if e != nil {
		fmt.Println("sholud bin err", e)
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	c.JSON(200, service.Modify())
}

func PostActionView(c *gin.Context) {
	var id = c.Query("id")
	pid, e := strconv.Atoi(id)

	if e != nil {
		c.JSON(200, rep.FatalResponseWithCode(gerr.ErrParamsInvlid))
		c.Abort()
		return
	}

	var srv = post.ActionPostService{Id: pid}

	c.JSON(200, srv.View())
}
