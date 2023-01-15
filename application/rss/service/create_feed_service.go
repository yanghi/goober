package service

import (
	"encoding/json"
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

	feed, r := getSrv.GetRawFromWeb()

	if r != nil {
		return goober.NewResponse().RawError(r).Msg("获取数据源失败,无法创建").Result()
	}

	authorJson, _ := json.Marshal(feed.Authors)

	dr, e := mysql.DB().Exec(
		"INSERT INTO gb_rss_feed (feed_link,title,description,link,authors,published,updated,version,feed_type,language) VALUES(?,?,?,?,?,?,?,?,?,?)",
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
	)

	if e != nil {
		var msg = "创建失败"
		if strings.Contains(e.Error(), "Duplicate entry") {
			msg = "订阅源已添加,请勿重复操作"
		}
		return goober.NewResponse().Label("添加订阅源").RawError(e).Msg(msg).AllowLog().Result()
	}
	lid, _ := dr.LastInsertId()

	e2 := s.insertFeedItems(lid, feed)

	if e2 != nil {
		mysql.DB().Exec("delete from gb_rss_feed where id=?", lid)

		return goober.NewResponse().AllowLog().RawError(e2).Msg("添加订阅源数据失败").Result()
	}

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
		sb.Values("", "", "", "", "", "", "", "", "", "")
		authors, _ := json.Marshal(it.Authors)
		cts, _ := json.Marshal(it.Categories)
		vals = append(vals, it.GUID, fid, it.Title, it.Description, it.Content, it.Link, it.PublishedParsed, it.Updated, authors, cts)
	}
	_, e2 := mysql.DB().Exec(sb.String(), vals...)

	if e2 != nil {
		return e2
	}

	return nil
}
