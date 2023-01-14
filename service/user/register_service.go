package user

import (
	"goober/application/mysql"
	"goober/auth"
	"goober/goober"
	"strings"
)

type RegisterService struct {
	Name     string `json:"name" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}

func (r *RegisterService) Register() *goober.ResponseResult {

	salt := auth.Salt()

	res, er := mysql.DB().Exec("insert into gb_user(name,password,salt) VALUES(?,?,?);", r.Name, auth.JoinSalt(r.Password, salt), salt)

	if er != nil {
		if strings.Contains(er.Error(), "Duplicate entry") {
			return goober.FailedResult(goober.ErrUnExpect, "用户名已存在")
		}
		return goober.FailedResult(goober.ErrDB, "注册失败")
	}

	id, er := res.LastInsertId()
	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "注册失败2")
	}

	return goober.OkResult(map[string]any{
		"user": map[string]any{
			"id":   id,
			"name": r.Name,
		},
	})
}
