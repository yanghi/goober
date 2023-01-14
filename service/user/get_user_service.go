package user

import (
	"goober/application/mysql"
	"goober/goober"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type GetUserService struct {
	Id     int   `json:"id"`
	IdList []int `json:"idList"`
}

type BaseUserMap map[string]interface{}

func (l *GetUserService) GetBaseInfoMap() (map[string]interface{}, error) {
	rows, er := mysql.DB().Query("select id,name,avatar_url as avatar from gb_user where id=?", l.Id)

	if er != nil {
		return nil, er
	}
	ms, er := goober.MysqlRowsToMap(rows)

	if er != nil {
		return nil, er
	}

	if len(ms) == 0 {
		return nil, goober.NewError().Msg("用户不存在")
	}

	return ms[0], nil
}
func (srv *GetUserService) GetBaseInfo() *goober.ResponseResult {
	info, e := srv.GetBaseInfoMap()

	if e != nil {
		return goober.WrongResult(e)
	}
	return goober.OkResult(info)
}
func (srv *GetUserService) GetUserBaseInfoListRaw() ([]map[string]interface{}, error) {
	var ids = ""

	for i, v := range srv.IdList {
		if i == 0 {
			ids += strconv.Itoa(v)
		} else {
			ids += "," + strconv.Itoa(v)
		}
	}

	rows, er := mysql.DB().Query("select id,name,avatar_url as avatar from gb_user where id in (?)", ids)

	if er != nil {
		return nil, er
	}
	ms, er := goober.MysqlRowsToMap(rows)

	if er != nil {
		return nil, er
	}

	return ms, nil
}
