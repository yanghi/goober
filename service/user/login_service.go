package user

import (
	"goober/application/mysql"
	"goober/auth"

	"goober/goober"

	_ "github.com/go-sql-driver/mysql"
)

type LoginService struct {
	Name     string `form:"name" json:"name" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

func (l *LoginService) Login() *goober.ResponseResult {
	r := l.ValidateParams()
	if r != nil {
		return r
	}

	rows, er := mysql.DB().Query("select * from gb_user where BINARY name=?", l.Name)

	if er != nil {
		return goober.WrongResult(er)
	}
	ms, er := goober.MysqlRowsToMap(rows)

	if er != nil {
		return goober.WrongResult(er)
	}

	if len(ms) == 0 {
		return goober.NewResponse().Msg("账号不存在").Code(goober.ErrNotExsit).Result()
	}

	user := ms[0]

	passwordCorrect := auth.Validate(l.Password, user["salt"].(string), user["password"].(string))

	if !passwordCorrect {
		return goober.FailedResult(goober.ErrParamsInvlid, "密码错误")
	}

	delete(user, "salt")
	delete(user, "password")

	return goober.OkResult(map[string]any{
		"user": user,
	})
}

func (l *LoginService) ValidateParams() *goober.ResponseResult {
	if l.Name == "" {
		return goober.FailedResult(goober.ErrParamsInvlid, "用户名格式错误")
	}

	if len(l.Password) < 6 {
		return goober.FailedResult(goober.ErrParamsInvlid, "密码长度不能小于6")
	}

	return nil
}
