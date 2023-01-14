package tag_service

import (
	"goober/application/mysql"
	"goober/goober"
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

func (srv *CreateTagService) Create() *goober.ResponseResult {
	res, er := mysql.DB().Exec("INSERT INTO gb_post_tag (name) VALUE(?);", srv.Name)
	if er != nil {
		if strings.Contains(er.Error(), "Duplicate entry") {
			return goober.FailedResult(goober.ErrUnExpect, "该标签名已存在,请使用其它名称")
		}

		return goober.FailedResult(goober.ErrDB, "创建标签失败")
	}
	id, er := res.LastInsertId()
	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "创建标签失败")
	}

	if er != nil {
		return goober.WrongResult(er)
	}

	return goober.OkResult(CreateTagResult{Name: srv.Name, Id: int(id)})
}
