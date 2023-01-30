package user

import (
	"goober/application/mysql"
	"goober/application/rss/service"
	"goober/goober"
	"strings"
)

type CreateFeedService struct {
	FeedId int `json:"feedId" form:"feedId"`
	UserId int
	Url    string `json:"url" form:"url"`
}

func (s *CreateFeedService) CreateWithUrl() *goober.ResponseResult {
	rs := service.CreateFeedService{Href: s.Url}

	r := rs.Create()

	if !r.Ok {
		return r
	}

	fid := r.Data.(map[string]interface{})["id"].(int)
	s.FeedId = fid

	return s.createWithFeedId()
}

func (s *CreateFeedService) CreateWithFeedId() *goober.ResponseResult {

	return s.createWithFeedId()
}
func (s *CreateFeedService) createWithFeedId() *goober.ResponseResult {
	_, e := mysql.DB().Exec("INSERT INTO gb_rss_user_feed (uid,feed_id) VALUES(?,?)", s.UserId, s.FeedId)

	if e != nil {
		if strings.Contains(e.Error(), "Duplicate entry") {
			return goober.FailedResult(goober.ErrExsited, "已添加,无重复操作")
		}

		return goober.ErrorLogResponse(e, "").Result()
	}

	return goober.OkResult(nil)
}
