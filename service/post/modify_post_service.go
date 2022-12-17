package post

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type ModifyPostService struct {
	Content     string `json:"content"`
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	Description string `json:"desription"`
}

func (srv *ModifyPostService) Modify() *rep.Response {

	var fields []string
	var values []any

	if srv.Content != "" {
		fields = append(fields, "author_id")
		values = append(values, srv.Content)
	}
	if srv.Title != "" {
		fields = append(fields, "title")
		values = append(values, srv.Title)
	}

	if len(fields) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "修改无有效字段")
	}
	rows, er := mysql.DB.Exec("DELETE FROM gb_post where author_id=? and id=?", values...)
	// rows, er := stm.Query(srv.Title, srv.Content, srv.AuthorId, srv.Description, tagQs.vals)
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "删除文章失败")
	}
	rowNum, _ := rows.RowsAffected()

	if rowNum == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在或无权限")
	}

	return rep.BuildOkResponse(nil)
}
