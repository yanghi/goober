package user

import (
	"goblog/auth"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"strings"
)

type RegisterService struct {
	Name     string `json:"name" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}

func (r *RegisterService) Register() *rep.Response {

	salt := auth.Salt()

	res, er := mysql.DB.Exec("insert into gb_user(name,password,salt) VALUES(?,?,?);", r.Name, auth.JoinSalt(r.Password, salt), salt)

	if er != nil {
		if strings.Contains(er.Error(), "Duplicate entry") {
			return rep.Build(nil, gerr.ErrUnExpect, "用户名已存在")
		}
		return rep.Build(nil, gerr.ErrDB, "注册失败")
	}

	id, er := res.LastInsertId()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "注册失败2")
	}

	return rep.BuildOkResponse(map[string]any{
		"user": map[string]any{
			"id":   id,
			"name": r.Name,
		},
	})
}
