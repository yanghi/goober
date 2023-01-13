package post

import (
	"fmt"
	"goober/database/mysql"
	gerr "goober/error"
	"goober/rep"
)

type DeletePostService struct {
	Id       int `json:"id"`
	AuthorId int `json:"authorId"`
}

func (srv *DeletePostService) DeleteByAuthor() *rep.Response {
	rows, er := mysql.DB.Exec("DELETE FROM gb_post where author_id=? and id=?", srv.AuthorId, srv.Id)

	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "删除文章失败")
	}
	rowNum, _ := rows.RowsAffected()

	fmt.Println("sss", srv.AuthorId, srv.Id)
	if rowNum == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在或无权限")
	}

	return rep.BuildOkResponse(map[string]any{
		"id": srv.Id,
	})
}

func (srv *DeletePostService) DeleteByAdmin() *rep.Response {
	return nil
}
