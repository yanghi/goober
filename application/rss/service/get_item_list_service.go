package service

import (
	"goober/application/mysql"
	"goober/goober"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

type GetItemListService struct {
	goober.PaginationQuery
	FeedId      int `json:"feedId" form:"feedId"`
	PPagination goober.Pagination
}

func (s *GetItemListService) Get() *goober.ResponseResult {

	sb := sqlbuilder.NewSelectBuilder()
	return s.get(sb)
}
func (s *GetItemListService) get(sb *sqlbuilder.SelectBuilder) *goober.ResponseResult {

	sb.Select("*").From("gb_rss_feed_items").OrderBy("id")
	s.PPagination.TouchSqlBuilder(sb)

	rs, e := mysql.DB().Query(sb.String())
	if e != nil {
		return goober.ErrorLogResponse(e, "获取订阅源数据列表").Result()
	}
	dt, e2 := goober.MysqlRowsToMap(rs)

	if e2 != nil {
		return goober.ErrorLogResponse(e2, "转换订阅源数据列表").Result()
	}
	totalRows, _ := mysql.DB().Query("select COUNT(*) total from gb_rss_feed_items where feed_id =?", s.FeedId)
	t, _ := goober.MysqlRowsToMap((totalRows))

	// total, _ := strconv.Atoi(t[0]["total"].(string))

	s.PPagination.Total(int(t[0]["total"].(int64)))

	return goober.OkResult(s.PPagination.ListMapResult(dt))
}

func (s *GetItemListService) GetToday() *goober.ResponseResult {
	sb := sqlbuilder.NewSelectBuilder()

	t := time.Now().Format("2006-01-02")

	// sb.Where(sb.Or(sb.Like("updated", t+"%"), sb.Like("updated", t+"%")))
	sb.Where(sb.Or("published LIKE \""+t+"%\"", "updated LIKE \""+t+"%\""))
	return s.get(sb)
}
