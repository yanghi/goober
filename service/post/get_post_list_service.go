package post

import (
	"fmt"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/model"
	"goblog/rep"
	"strconv"
)

type GetPostListService struct {
	PaginationParams model.PaginationParams
	AuthorId         int `json:"authorId"`
}

func (srv *GetPostListService) GetByAuthor() *rep.Response {

	// select * from gb_post where author_id=7 limit 0,10;SELECT FOUND_ROWS() as total;

	pg := model.Pagination{}
	p := pg.Params(srv.PaginationParams)

	rows, e := mysql.DB.Query("select * from gb_post where author_id=?;", srv.AuthorId, "limit ?,?;", pg.Start(), p.Size)

	if e != nil {
		return rep.FatalResponseWithCode(gerr.ErrDB)
	}
	ms, er := mysql.RowsToMap(rows)

	if len(ms) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在")
	}

	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取文章失败,数据转换失败")
	}

	return rep.BuildOkResponse(ms)
}
func (srv *GetPostListService) Get() *rep.Response {

	pg := model.Pagination{}
	p := pg.Params(srv.PaginationParams)

	qbuilder := mysql.QueryBuilder{}

	rows, e := mysql.DB.Query("select * from gb_post "+qbuilder.BuildOrderBy([]string{"id"}, p.Order)+" limit ?,?", pg.Start(), p.Size)
	totalRows, _ := mysql.DB.Query("select COUNT(*) total from gb_post;")
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

	return rep.BuildOkResponse(map[string]interface{}{
		"total": total,
		"page":  p.Page,
		"size":  p.Size,
		"order": p.Order,
		"list":  ms,
	})
}
