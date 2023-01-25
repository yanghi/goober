package service

import (
	"goober/application/mysql"
	"goober/goober"
)

type GetFeedListService struct{}

func (s *GetFeedListService) GetAll() *goober.ResponseResult {
	r, e := mysql.DB().Query(
		"SELECT a.*,b.count from gb_rss_feed a INNER JOIN (SELECT feed_id, COUNT(*) as count FROM gb_rss_feed_items GROUP BY feed_id) b ON a.id = b.feed_id;")

	if e != nil {
		return goober.WrongResult(e)
	}

	dt, e2 := goober.MysqlRowsToMap(r)

	if e2 != nil {
		return goober.ErrorLogResponse(e2, "").Result()
	}

	return goober.OkResult(dt)
}
