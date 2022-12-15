package user

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type RegisterService struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *RegisterService) Register() *rep.Response {

	stm, _ := mysql.DB.Prepare("insert IGNORE into gb_user(name,password) VALUES(?,MD5(?));")

	res, er := stm.Exec(r.Name, r.Password)

	if er != nil {
		return rep.Build(nil, gerr.ErrDB, "注册失败")
	}

	id, er := res.LastInsertId()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "注册失败")
	}

	return rep.BuildOkResponse(map[string]any{
		"user": map[string]any{
			"id":   id,
			"name": r.Name,
		},
	})
}
