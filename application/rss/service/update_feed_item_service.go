package service

import (
	"fmt"
	"goober/application/mysql"
	"goober/goober"

	"github.com/mmcdole/gofeed"
)

type UpdateFeedItemsService struct {
}
type UpdateSingleFeedItemsResult struct {
	Success     bool   `json:"success"`
	UpdateCount int    `json:"updateCount"`
	Reason      string `json:"reason"`
	Msg         string `json:"msg"`
	Href        string `json:"href"`
	// feed id
	Id int64 `json:"id"`
}
type UpdateFeedResult struct {
	SuccessCount int                           `json:"successCount"`
	TotalCount   int                           `json:"total"`
	FailedCount  int                           `json:"failedCount"`
	ValidCount   int                           `json:"validCount"`
	Result       []UpdateSingleFeedItemsResult `json:"result"`
}

func (s *UpdateFeedItemsService) UpdateSinge(href string) UpdateSingleFeedItemsResult {
	var gs = GetFeedService{Href: href}
	res := UpdateSingleFeedItemsResult{Href: href}

	r, e := mysql.DB().Query("select id from gb_rss_feed where feed_link=?", href)
	fmt.Println("fs", e)

	if e != nil {
		res.Reason = e.Error()
		return res
	}

	idRows, _ := goober.MysqlRowsToMap(r)

	if len(idRows) == 0 {
		res.Reason = "不存在"
		return res
	}
	feed, e := gs.GetRawFromWeb()

	if e != nil {
		res.Msg = "获取数据源失败"
		res.Reason = e.Error()
		return res
	}

	fid := idRows[0]["id"].(int64)
	res.Id = fid

	if len(feed.Items) > 0 {
		r2, e3 := mysql.DB().Query("select id,published from gb_rss_feed_items where feed_id=?  order by id DESC limit 0,1", fid)

		if e3 != nil {
			res.Reason = e3.Error()
			return res
		}

		lastIdm, _ := goober.MysqlRowsToMap(r2)

		if len(lastIdm) > 0 {
			last := lastIdm[0]["published"]
			lastTime := last.(string)
			latest := []*gofeed.Item{}

			for _, i := range feed.Items {

				if i.PublishedParsed.Format("2006-01-02 15:04:05") > lastTime {
					latest = append(latest, i)
				}
			}
			feed.Items = latest
		}

		if len(feed.Items) > 0 {
			var cs = CreateFeedService{}
			e4 := cs.insertFeedItems(fid, feed)
			fmt.Println("insert feed item error", e4)
		} else {
			res.Msg = "已经是最新的"
		}
	}

	res.Success = true
	res.UpdateCount = len(feed.Items)
	return res
}
func (s *UpdateFeedItemsService) UpdateAll() *goober.ResponseResult {
	fr, e := mysql.DB().Query("select feed_link from gb_rss_feed")

	if e != nil {
		return goober.NewResponse().RawError(e).AllowLog().FailedResult()
	}

	fsmap, _ := goober.MysqlRowsToMap(fr)
	var res = UpdateFeedResult{}
	res.TotalCount = len(fsmap)

	if len(fsmap) > 0 {
		for _, feed := range fsmap {
			fl := feed["feed_link"].(string)
			upRes := s.UpdateSinge(fl)
			res.Result = append(res.Result, upRes)

			if upRes.Success {
				res.SuccessCount++
				res.ValidCount += upRes.UpdateCount
			} else {
				res.FailedCount++
			}
		}
	}

	return goober.OkResult(res)
}
