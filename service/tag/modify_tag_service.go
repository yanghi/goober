package tag_service

import (
	"goober/application/mysql"
	"goober/goober"
)

type ModifyTagService struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (srv *ModifyTagService) Modify() *goober.ResponseResult {
	res, er := mysql.DB().Exec("UPDATE gb_post_tag SET name=? where id=?", srv.Name, srv.Id)
	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "修改标签失败")
	}
	count, er := res.RowsAffected()
	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "修改标签失败")
	}

	if count == 0 {
		return goober.FailedResult(goober.ErrUnExpect, "修改标签失败,标签不存在或无权限")
	}

	return goober.OkResult(nil)
}
