package post

import (
	"goober/application/mysql"
	"goober/goober"
)

type ActionPostService struct {
	Id int `json:"id"`
}

func (srv *ActionPostService) View() *goober.ResponseResult {
	_, er := mysql.DB().Query("update gb_post set view = view +1 where id=?", srv.Id)

	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "view error")
	}

	return goober.OkResult(nil)
}
