package tag_service

import (
	"goober/database/mysql"
	gerr "goober/error"
	"goober/rep"
)

type DeleteTagService struct {
	Id int `json:"id"`
}

type DeleteTagResult struct {
	Id int `json:"id"`
}

func (srv *DeleteTagService) Del() *rep.Response {
	res, er := mysql.DB.Exec("DELETE FROM gb_post_tag where id=?", srv.Id)
	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "删除失败")
	}
	count, _ := res.RowsAffected()

	if count == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "删除失败,标签不存在或无权限")
	}

	return rep.BuildOkResponse(DeleteTagResult{Id: srv.Id})
}
