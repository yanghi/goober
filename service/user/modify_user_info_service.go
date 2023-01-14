package user

import (
	"fmt"
	"goober/application/mysql"
	"goober/goober"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type ModifyUserInfoService struct {
	Id     int    `json:"id"`
	Name   string `json:"name" binding:"min=2,max=30"`
	Avatar string `json:"avatar"`
}

func (s *ModifyUserInfoService) Modify() *goober.ResponseResult {
	var fields []string
	var values []any

	if s.Avatar != "" {
		fields = append(fields, "avatar_url=?")
		values = append(values, s.Avatar)
	}
	if s.Name != "" {
		fields = append(fields, "name=?")
		values = append(values, s.Name)
	}
	if len(fields) == 0 {
		return goober.FailedResult(goober.ErrUnExpect, "修改无有效字段")
	}
	values = append(values, s.Id)

	_, e := mysql.DB().Exec("UPDATE gb_user SET "+strings.Join(fields, ",")+" WHERE id=?", values...)

	if e != nil {
		fmt.Println("修改错误", e)
		return goober.FailedResult(goober.ErrUnExpect, "修改失败")
	}

	return goober.OkResult(nil)
}
