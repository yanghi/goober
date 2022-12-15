package user

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"

	_ "github.com/go-sql-driver/mysql"
)

type LoginService struct {
	Name     string `form:"name" json:"name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

func (l *LoginService) Login() *rep.Response {
	r := l.ValidateParams()
	if r != nil {
		return r
	}

	stm, _ := mysql.DB.Prepare("select id,name,time from gb_user where BINARY name=? and password=MD5(?);")

	rows, er := stm.Query(l.Name, l.Password)

	if er != nil {
		return rep.BuildFatalResponse(er)
	}
	ms, er := mysql.RowsToMap(rows)

	if er != nil {
		return rep.BuildFatalResponse(er)
	}

	if len(ms) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "账号或密码错误")
	}

	return rep.BuildOkResponse(map[string]any{
		"user": ms[0],
	})
}

func (l *LoginService) ValidateParams() *rep.Response {
	if l.Name == "" {
		return rep.Build(nil, gerr.ErrParamsInvlid, "用户名格式错误")
	}

	if len(l.Password) < 6 {
		return rep.Build(nil, gerr.ErrParamsInvlid, "密码长度不能小于6")
	}

	return nil
}
