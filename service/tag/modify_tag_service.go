package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type ModifyTagService struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (srv *ModifyTagService) Modify() *rep.Response {
	res, er := mysql.DB.Exec("UPDATE gb_post_tag SET name=? where id=?", srv.Name, srv.Id)
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
