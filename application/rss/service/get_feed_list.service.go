package service

import (
	"goober/application/mysql"
	"goober/goober"
)

type GetFeedListService struct{}

func (s *GetFeedListService) GetAll() *goober.ResponseResult {
	r, e := mysql.DB().Query("select * from gb_rss_feed")

	if e != nil {
		return goober.WrongResult(e)
	}

	dt, e2 := goober.MysqlRowsToMap(r)

	if e2 != nil {
		return goober.ErrorLogResponse(e2, "").Result()
	}

	return goober.OkResult(dt)
}
