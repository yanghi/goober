package service

import (
	"encoding/json"
	"fmt"
	"goober/application/mysql"
	"goober/goober"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/mmcdole/gofeed"
)

type CreateFeedService struct {
	Href string `json:"href"`
}

func (s *CreateFeedService) Create() *goober.ResponseResult {
	var getSrv = GetFeedService{Href: s.Href}

	feed, r := getSrv.GetRawWithUrl()

	if r != nil {
		return goober.ErrorLogResponse(r, "[rss]").Msg("获取数据源失败,无法创建").Result()
	}

	authorJson, _ := json.Marshal(feed.Authors)

	tx, e0 := mysql.DB().Begin()

	if e0 != nil {
		tx.Rollback()
		return goober.ErrorLogResponse(e0, "[rss]").Result()
	}

	dr, e := tx.Exec(
		"INSERT INTO gb_rss_feed (feed_link,title,description,link,authors,published,updated,version,feed_type,language,rsshub) VALUES(?,?,?,?,?,?,?,?,?,?,?)",
		feed.FeedLink,
		feed.Title,
		feed.Description,
		feed.Link,
		authorJson,
		feed.Published,
		feed.UpdatedParsed,
		feed.FeedVersion,
		feed.FeedType,
		feed.Language,
		getSrv.IsHubFeedUrl(),
	)

	if e != nil {
		tx.Rollback()
		var msg = "创建失败"
		if strings.Contains(e.Error(), "Duplicate entry") {
			msg = "订阅源已添加,请勿重复操作"
		}
		fmt.Println("psh", feed.Published, "--", feed.PublishedParsed)
		return goober.NewResponse().Label("添加订阅源").RawError(e).Msg(msg).AllowLog().Result()
	}
	lid, _ := dr.LastInsertId()

	e2 := s.insertFeedItems(lid, feed)

	if e2 != nil {
		tx.Rollback()
		return goober.NewResponse().AllowLog().RawError(e2).Msg("添加订阅源数据失败").Result()
	}

	tx.Commit()

	return goober.OkResult(map[string]interface{}{
		"id":   lid,
		"feed": feed,
	})
}

func (s *CreateFeedService) insertFeedItems(fid int64, feed *gofeed.Feed) error {
	sb := sqlbuilder.InsertInto("gb_rss_feed_items").
		Cols("guid", "feed_id", "title", "description", "content", "link", "published", "updated", "authors", "categories")

	vals := []any{}

	for l := len(feed.Items) - 1; l >= 0; l-- {
		it := feed.Items[l]

		// 部分转换后的时间是无效的
		if it.PublishedParsed == nil {
			goober.Logger().Error("[rss] rss数据源时间无效", it.Title, it.Link, feed.FeedLink)
			continue
		}

		sb.Values("", "", "", "", "", "", "", "", "", "")
		authors, _ := json.Marshal(it.Authors)
		cts, _ := json.Marshal(it.Categories)

		vals = append(vals, it.GUID, fid, it.Title, it.Description, it.Content, it.Link, it.PublishedParsed.Format("2006-01-02 15:04:05"), it.Updated, authors, cts)
	}
	_, e2 := mysql.DB().Exec(sb.String(), vals...)

	if e2 != nil {
		return e2
	}

	return nil
}
