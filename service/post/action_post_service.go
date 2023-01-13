package post

import (
	"goober/database/mysql"
	gerr "goober/error"
	"goober/rep"
)

type ActionPostService struct {
	Id int `json:"id"`
}

func (srv *ActionPostService) View() *rep.Response {
	_, er := mysql.DB.Query("update gb_post set view = view +1 where id=?", srv.Id)

	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "view error")
	}

	return rep.BuildOkResponse(nil)
}
