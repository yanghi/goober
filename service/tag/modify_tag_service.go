package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type ModifyTagService struct {
	Text string `json:"text"`
	Id   int    `json:"id"`
}

func (srv *ModifyTagService) Modify() *rep.Response {
	res, er := mysql.DB.Exec("UPDATE gb_post_tag SET text=? where id=?", srv.Text, srv.Id)
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "修改标签失败")
	}
	count, er := res.RowsAffected()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "修改标签失败")
	}

	if count == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "修改标签失败,标签不存在或无权限")
	}

	return rep.BuildOkResponse(nil)
}
