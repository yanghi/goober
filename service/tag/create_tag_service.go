package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"strings"
)

type CreateTagService struct {
	Name string `json:"name"`
	// KindId int    `json:"kindId"`
}
type CreateTagResult struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	// KindId int    `json:"kindId"`
}

func (srv *CreateTagService) Create() *rep.Response {
	res, er := mysql.DB.Exec("INSERT INTO gb_post_tag (name) VALUE(?);", srv.Name)
	if er != nil {
		if strings.Contains(er.Error(), "Duplicate entry") {
			return rep.Build(nil, gerr.ErrUnExpect, "该标签名已存在,请使用其它名称")
		}

		return rep.Build(nil, gerr.ErrDB, "创建标签失败")
	}
	id, er := res.LastInsertId()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "创建标签失败")
	}

	if er != nil {
		return rep.BuildFatalResponse(er)
	}

	return rep.BuildOkResponse(CreateTagResult{Name: srv.Name, Id: int(id)})
}
