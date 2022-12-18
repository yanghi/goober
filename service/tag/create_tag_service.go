package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type CreateTagService struct {
	Text string `json:"text"`
	// KindId int    `json:"kindId"`
}
type CreateTagResult struct {
	Text string `json:"text"`
	Id   int    `json:"id"`
	// KindId int    `json:"kindId"`
}

func (srv *CreateTagService) Create() *rep.Response {
	res, er := mysql.DB.Exec("INSERT INTO gb_post_tag (text) VALUE(?);", srv.Text)
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "创建标签失败")
	}
	id, er := res.LastInsertId()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "创建标签失败")
	}

	if er != nil {
		return rep.BuildFatalResponse(er)
	}

	return rep.BuildOkResponse(CreateTagResult{Text: srv.Text, Id: int(id)})
}
