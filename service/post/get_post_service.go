package post

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"goblog/service/user"
)

type GetPostService struct {
	Id int `json:"id"`
}

func (srv *GetPostService) Get() *rep.Response {

	rows, er := mysql.DB.Query("SELECT * FROM gb_post where id=?", srv.Id)
	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取文章失败")
	}

	ms, er := mysql.RowsToMap(rows)

	if len(ms) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在")
	}

	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取文章失败,数据转换失败")
	}
	post := ms[0]

	var usrv = user.GetUserService{Id: int(post["author_id"].(int64))}

	var u, _ = usrv.GetBaseInfoMap()
	post["author"] = u
	return rep.BuildOkResponse(post)
}
