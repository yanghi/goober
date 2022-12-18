package tag_service

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"strconv"
	"strings"
)

type GetTagListService struct {
	IdList []int `json:"idList" form:"idList"`
}

func (srv *GetTagListService) Get() *rep.Response {
	return srv.get("")
}
func (srv *GetTagListService) GetByIdList() *rep.Response {
	vals := []string{}

	for _, v := range srv.IdList {
		vals = append(vals, strconv.Itoa(v))
	}

	return srv.get("where id in (" + strings.Join(vals, ",") + ")")
}

func (srv *GetTagListService) get(queryPart string) *rep.Response {
	// querys:=[]string{"SELECT * FROM gb_post_tag "}

	rows, e := mysql.DB.Query("SELECT * FROM gb_post_tag " + queryPart)

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
