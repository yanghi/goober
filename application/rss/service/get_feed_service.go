package service

import (
	"context"
	"goober/goober"
	"time"

	"github.com/mmcdole/gofeed"
)

type GetFeedService struct {
	Href string `json:"href" form:"href"`
}

func (s *GetFeedService) GetRawFromWeb() (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	feed, err := fp.ParseURLWithContext(s.Href, ctx)

	if err != nil {
		return nil, err
	}
	return feed, nil
}
func (s *GetFeedService) GetFromWeb() *goober.ResponseResult {
	feed, err := s.GetRawFromWeb()

	if err != nil {
		return goober.ErrorLogResponse(err, "获取网络rss").Result()
	}

	return goober.OkResult(feed)
}
