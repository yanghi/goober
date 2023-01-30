package service

import (
	"context"
	"goober/application/rss"
	"goober/goober"
	"time"

	"github.com/mmcdole/gofeed"
)

type GetFeedService struct {
	Href    string `json:"href" form:"href"`
	_isHub  bool
	_hubUrl string
	_a      string
}

func (s *GetFeedService) GetRawWithUrl() (*gofeed.Feed, error) {

	fp := gofeed.NewParser()
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	feed, err := fp.ParseURLWithContext(rss.GetRealFeedUrl(s.Href), ctx)

	if err != nil {
		return nil, err
	}

	formatFeed(feed)

	return feed, nil
}
func (s *GetFeedService) GetWithUrl() *goober.ResponseResult {

	feed, err := s.GetRawWithUrl()

	if err != nil {
		return goober.ErrorLogResponse(err, "获取rss").Result()
	}

	return goober.OkResult(feed)
}

func (s *GetFeedService) GetFromWeb() *goober.ResponseResult {
	if s.IsHubFeedUrl() {
		return goober.FailedResult(goober.ErrUnExpect, "获取失败")
	}
	return s.GetWithUrl()
}

func (s *GetFeedService) IsHubFeedUrl() bool {
	s._isHub = rss.IsHubFeedUrl(s.Href)
	return s._isHub
}

func formatFeed(feed *gofeed.Feed) {
	if len(feed.Items) != 0 {
		for _, it := range feed.Items {

			if it.Published == "" && it.Updated != "" {
				it.Published = it.Updated
				t, e := rss.ParseDate(it.Published)
				if e != nil {
					it.PublishedParsed = &t
				}
			}
			if it.Updated == "" {
				it.Updated = it.Published
			}

			if it.Updated == "" && it.Published == "" {
				goober.Logger().Error("[goober rss]: item缺少时间信息")
			}

			if it.GUID == "" {
				it.GUID = it.Link
			}

		}

		if feed.Published == "" {
			feed.Published = feed.Items[0].Published
			feed.PublishedParsed = feed.Items[0].PublishedParsed
		}
	}
}
