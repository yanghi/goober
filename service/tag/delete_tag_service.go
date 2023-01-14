package tag_service

import (
	"goober/application/mysql"
	"goober/goober"
)

type DeleteTagService struct {
	Id int `json:"id"`
}

type DeleteTagResult struct {
	Id int `json:"id"`
}

func (srv *DeleteTagService) Del() *goober.ResponseResult {
	res, er := mysql.DB().Exec("DELETE FROM gb_post_tag where id=?", srv.Id)
	if er != nil {
		return goober.FailedResult(goober.ErrDB, "删除失败")
	}
	count, _ := res.RowsAffected()

	if count == 0 {
		return goober.FailedResult(goober.ErrUnExpect, "删除失败,标签不存在或无权限")
	}

	return goober.OkResult(DeleteTagResult{Id: srv.Id})
}
