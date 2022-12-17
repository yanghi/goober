package user

import (
	"goblog/database/mysql"
	gerr "goblog/error"

	_ "github.com/go-sql-driver/mysql"
)

type GetUserService struct {
	Id int `json:"id"`
}

type BaseUserMap map[string]interface{}

func (l *GetUserService) GetBaseInfoMap() (map[string]interface{}, error) {
	rows, er := mysql.DB.Query("select id,name from gb_user where id=?", l.Id)

	if er != nil {
		return nil, er
	}
	ms, er := mysql.RowsToMap(rows)

	if er != nil {
		return nil, er
	}

	if len(ms) == 0 {
		return nil, gerr.New("用户不存在")
	}

	return ms[0], nil
}
