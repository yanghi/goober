package post

import (
	"encoding/json"
	"fmt"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/model"
	"goblog/model/post"
	"goblog/rep"
	"goblog/serializer"
	tag_service "goblog/service/tag"
	user_service "goblog/service/user"
	"strconv"

	"github.com/huandu/go-sqlbuilder"
)

type GetPostListService struct {
	PaginationParams model.PaginationParams
	AuthorId         int `json:"authorId"`
	Statu            post.PostStatu
}

func (srv *GetPostListService) GetByAuthor() *rep.Response {
	where := "author_id=" + strconv.Itoa(srv.AuthorId)

	sql := sqlbuilder.Select("*").From("gb_post").Where(where)

	if srv.Statu != post.PostStatuNotExsit {
		w2 := "statu=" + strconv.Itoa(int(post.PostStatuPublic))
		sql.Where(w2)

		where += " and " + w2

	}

	return srv.get(sql, where)
}

func (srv *GetPostListService) Get() *rep.Response {
	where := "statu=" + strconv.Itoa(int(post.PostStatuPublic))
	sql := sqlbuilder.Select("*").From("gb_post").Where(where)

	return srv.get(sql, where)
}

func (srv *GetPostListService) get(builder *sqlbuilder.SelectBuilder, where string) *rep.Response {

	if where != "" {
		where = " where " + where
	}

	pg := model.Pagination{}
	p := pg.Params(srv.PaginationParams)

	builder.Limit(p.Size).Offset(pg.Start()).OrderBy("id")

	if p.Order == "ASC" {
		builder.Asc()
	} else {
		builder.Desc()
	}

	rows, e := mysql.DB.Query(builder.String())
	totalRows, _ := mysql.DB.Query("select COUNT(*) total from gb_post" + where)
	if e != nil {
		fmt.Println("get postlist:", e)
		return rep.FatalResponseWithCode(gerr.ErrDB)
	}
	ms, er := mysql.RowsToMap(rows)

	t, _ := mysql.RowsToMap((totalRows))
	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取文章失败,数据转换失败")
	}

	total, _ := strconv.Atoi(t[0]["total"].(string))

	tagIdMap := map[int]int{}
	authorIdMap := map[int]int{}

	for _, post := range ms {

		// 格式化
		id, _ := strconv.Atoi(post["id"].(string))
		post["id"] = id

		view, _ := strconv.Atoi(post["view"].(string))
		post["view"] = view
		statu, _ := strconv.Atoi(post["statu"].(string))
		post["statu"] = statu
		author_id, _ := strconv.Atoi(post["author_id"].(string))
		post["authorId"] = author_id

		tagIds := []int{}

		authorIdMap[author_id] = author_id

		if post["tag"] != nil {
			json.Unmarshal([]byte(post["tag"].(string)), &tagIds)
		}

		post["tag"] = tagIds
		post["tagList"] = []map[string]any{}

		delete(post, "author_id")

		if post["description"].(string) == "" {
			post["description"] = serializer.Post.ExtractMarkdownDescription(post["content"].(string))
		}
		delete(post, "content")

		for _, id := range tagIds {
			tagIdMap[id] = id
		}
	}
	allTagIds := []int{}
	allAuthorIds := []int{}

	for _, v := range tagIdMap {
		allTagIds = append(allTagIds, v)
	}

	for _, v := range authorIdMap {
		allAuthorIds = append(allAuthorIds, v)
	}

	var tsrv = tag_service.GetTagListService{IdList: allTagIds}
	tagRes := tsrv.GetByIdList()

	if tagRes.Ok {
		var tagList = tagRes.Data.(map[string]any)["list"]
		list, ok := tagList.([]map[string]any)
		if ok {
			tabObjMap := map[int]any{}
			for _, v := range list {
				tabObjMap[v["id"].(int)] = v
			}

			for _, post := range ms {
				tagIds := post["tag"].([]int)
				tagList := []any{}

				for _, tid := range tagIds {
					tagObj, ok := tabObjMap[tid]
					if ok {
						tagList = append(tagList, tagObj)
					}
				}
				post["tagList"] = tagList
			}

		}
	}

	var usrv = user_service.GetUserService{IdList: allAuthorIds}

	userList, e := usrv.GetUserBaseInfoListRaw()

	if e == nil {
		userMap := map[int]any{}

		for _, v := range userList {
			userMap[int(v["id"].(int64))] = v
		}

		for _, v := range ms {
			user, ok := userMap[v["authorId"].(int)]

			if ok {
				v["author"] = user
			}
		}
	}

	totalPages := total / p.Size

	if total%p.Size != 0 {
		totalPages++
	}

	return rep.BuildOkResponse(map[string]interface{}{
		"total":      total,
		"page":       p.Page,
		"size":       p.Size,
		"order":      p.Order,
		"list":       ms,
		"totalPages": totalPages,
	})
}
