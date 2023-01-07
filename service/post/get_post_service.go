package post

import (
	"encoding/json"
	"fmt"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/model/post"
	"goblog/rep"
	tag_service "goblog/service/tag"
	"goblog/service/user"
	"strconv"

	"github.com/huandu/go-sqlbuilder"
)

type GetPostService struct {
	Id     int `json:"id"`
	UserId int
}

func (srv *GetPostService) GetByUser() *rep.Response {
	bd := sqlbuilder.Select("*").From("gb_post")
	bd.Where("id="+strconv.Itoa(srv.Id), bd.And("author_id="+strconv.Itoa(srv.UserId)+" OR statu="+post.PostStatuToStr(post.PostStatuPublic)))

	return srv.get(bd)
}

func (srv *GetPostService) Get() *rep.Response {
	bd := sqlbuilder.Select("*").From("gb_post")

	bd.Where("id="+strconv.Itoa(srv.Id), bd.And("statu="+post.PostStatuToStr(post.PostStatuPublic)))

	return srv.get(bd)
}

func (srv *GetPostService) get(sqlbd *sqlbuilder.SelectBuilder) *rep.Response {

	rows, er := mysql.DB.Query(sqlbd.String())
	if er != nil {
		fmt.Println("get post error: ", er, sqlbd.String())
		return rep.Build(nil, gerr.ErrDB, "获取文章失败")
	}

	ms, er := mysql.RowsToMap(rows)

	if len(ms) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在或者无权限")
	}

	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取文章失败,数据转换失败")
	}
	post := ms[0]

	// 格式化
	id, _ := strconv.Atoi(post["id"].(string))
	post["id"] = id

	view, _ := strconv.Atoi(post["view"].(string))
	post["view"] = view
	statu, _ := strconv.Atoi(post["statu"].(string))
	post["statu"] = statu
	author_id, e := strconv.Atoi(post["author_id"].(string))

	post["authorId"] = author_id
	delete(post, "author_id")

	if e == nil {
		var usrv = user.GetUserService{Id: author_id}

		var u, _ = usrv.GetBaseInfoMap()
		post["author"] = u
	} else {
		post["author"] = nil
	}

	tagIds := []int{}

	if post["tag"] != nil {
		json.Unmarshal([]byte(post["tag"].(string)), &tagIds)
	} else {
		post["tagList"] = []any{}
	}

	post["tag"] = tagIds

	if len(tagIds) > 0 {
		var tsrv = tag_service.GetTagListService{IdList: tagIds}

		tagRes := tsrv.GetByIdList()

		if tagRes.Ok {
			var tagList = tagRes.Data.(map[string]any)["list"]
			_, e := tagList.([]map[string]any)
			if e {
				post["tagList"] = tagList

			} else {

				post["tagList"] = []map[string]any{}
			}
		}
	}

	return rep.BuildOkResponse(post)
}
