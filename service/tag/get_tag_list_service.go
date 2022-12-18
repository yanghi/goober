package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"strconv"
)

type GetTagListService struct {
}

func (srv *GetTagListService) Get() *rep.Response {
	rows, e := mysql.DB.Query("SELECT * FROM gb_post_tag")

	if e != nil {
		return rep.FatalResponseWithCode(gerr.ErrDB)
	}
	ms, er := mysql.RowsToMap(rows)

	for _, v := range ms {
		// 为什么是string?
		id, _ := strconv.Atoi(v["id"].(string))
		v["id"] = id
	}
	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "获取标签失败,数据转换失败")
	}
	return rep.BuildOkResponse(map[string]interface{}{
		"list": ms,
	})
}
