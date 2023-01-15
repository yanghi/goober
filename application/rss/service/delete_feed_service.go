package service

import (
	"goober/application/mysql"
	"goober/goober"
)

type DeleteFeedService struct {
	Id int `json:"id" form:"id"`
}

func (s *DeleteFeedService) Delete() *goober.ResponseResult {
	tx, e := mysql.DB().Begin()

	if e != nil {
		tx.Rollback()
		return goober.ErrorLogResponse(e, "删除feed").Result()
	}

	_, e2 := tx.Exec("delete from gb_rss_feed where id=?", s.Id)

	if e2 != nil {
		tx.Rollback()
		return goober.ErrorLogResponse(e2, "删除feed").Result()
	}
	_, e3 := tx.Exec("delete from gb_rss_feed_items where feed_id=?", s.Id)

	if e3 != nil {
		tx.Rollback()
		return goober.ErrorLogResponse(e3, "删除feed").Result()
	}

	tx.Commit()

	return goober.OkResult(nil)
}
