package user

import (
	"goblog/auth"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"

	_ "github.com/go-sql-driver/mysql"
)

type LoginService struct {
	Name     string `form:"name" json:"name" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

func (l *LoginService) Login() *rep.Response {
	r := l.ValidateParams()
	if r != nil {
		return r
	}

	rows, er := mysql.DB.Query("select * from gb_user where BINARY name=?", l.Name)

	if er != nil {
		return rep.BuildFatalResponse(er)
	}
	ms, er := mysql.RowsToMap(rows)

	if er != nil {
		return rep.BuildFatalResponse(er)
	}

	if len(ms) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "账号不存在")
	}

	user := ms[0]

	passwordCorrect := auth.Validate(l.Password, user["salt"].(string), user["password"].(string))

	if !passwordCorrect {
		return rep.Build(nil, gerr.ErrParamsInvlid, "密码错误")
	}

	delete(user, "salt")
	delete(user, "password")

	return rep.BuildOkResponse(map[string]any{
		"user": user,
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
