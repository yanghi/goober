package user

import (
	"goober/application/mysql"
	"goober/goober"
)

type GetFeedListService struct {
	UserId int
}

func (s *GetFeedListService) GetAll() *goober.ResponseResult {

	// sb := sqlbuilder.Select("SELECT a.*").From("gb_rss_feed").JoinWithOption(sqlbuilder.InnerJoin, "(select * FROM gb_rss_user_feed WHERE uid = ?)", "b.feed_id = a.id;")
	rs, e := mysql.DB().Query("SELECT a.* FROM gb_rss_feed a INNER JOIN (select * FROM gb_rss_user_feed WHERE uid = ?) b ON b.feed_id = a.id;", s.UserId)

	if e != nil {
		return goober.ErrorLogResponse(e, "获取用户rss").Result()
	}
	dt, _ := goober.MysqlRowsToMap(rs)

	return goober.OkResult(dt)
}
