package post

import (
	"fmt"
	"goober/application/mysql"
	"goober/goober"
)

type DeletePostService struct {
	Id       int `json:"id"`
	AuthorId int `json:"authorId"`
}

func (srv *DeletePostService) DeleteByAuthor() *goober.ResponseResult {
	rows, er := mysql.DB().Exec("DELETE FROM gb_post where author_id=? and id=?", srv.AuthorId, srv.Id)

	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "删除文章失败")
	}
	rowNum, _ := rows.RowsAffected()

	fmt.Println("sss", srv.AuthorId, srv.Id)
	if rowNum == 0 {
		return goober.FailedResult(goober.ErrUnExpect, "文章不存在或无权限")
	}

	return goober.OkResult(map[string]any{
		"id": srv.Id,
	})
}

func (srv *DeletePostService) DeleteByAdmin() *goober.ResponseResult {
	return nil
}
