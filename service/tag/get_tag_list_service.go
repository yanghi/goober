package tag_service

import (
	"goober/application/mysql"
	"goober/goober"
	"strconv"
	"strings"
)

type GetTagListService struct {
	IdList []int `json:"idList" form:"idList"`
}

func (srv *GetTagListService) Get() *goober.ResponseResult {
	return srv.get("")
}
func (srv *GetTagListService) GetByIdList() *goober.ResponseResult {
	vals := []string{}

	for _, v := range srv.IdList {
		vals = append(vals, strconv.Itoa(v))
	}

	return srv.get("where id in (" + strings.Join(vals, ",") + ")")
}

func (srv *GetTagListService) get(queryPart string) *goober.ResponseResult {
	// querys:=[]string{"SELECT * FROM gb_post_tag "}

	rows, e := mysql.DB().Query("SELECT * FROM gb_post_tag " + queryPart)

	if e != nil {
		return goober.FailedResult(goober.ErrDB, "")
	}
	ms, er := goober.MysqlRowsToMap(rows)

	for _, v := range ms {
		// 为什么是string?
		id, _ := strconv.Atoi(v["id"].(string))
		v["id"] = id
	}
	if er != nil {
		return goober.FailedResult(goober.ErrDB, "获取标签失败,数据转换失败")
	}
	return goober.OkResult(map[string]interface{}{
		"list": ms,
	})
}
