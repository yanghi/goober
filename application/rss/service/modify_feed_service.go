package service

import (
	"goober/application/mysql"
	"goober/goober"

	"github.com/huandu/go-sqlbuilder"
)

type ModifyFeedService struct {
	Id    int    `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
}

func (s *ModifyFeedService) Modify() *goober.ResponseResult {
	ub := sqlbuilder.NewUpdateBuilder().Update("gb_rss_feed")
	ub.Where("id=" + ub.Args.Add(s.Id))

	if s.Title != "" {
		ub.Set("title=" + ub.Args.Add(s.Title))
	}

	sql, args := ub.Build()

	if len(args) == 1 {
		return goober.FailedResult(goober.ErrParamsInvlid, "至少需要一个有效修改字段")
	}
	r, e := mysql.DB().Exec(sql, args...)

	if e != nil {
		return goober.ErrorLogResponse(e, "修改订阅主题失败").Result()
	}
	ef, _ := r.RowsAffected()

	if ef == 0 {
		return goober.FailedResult(goober.ErrUnExpect, "修改失败")
	}

	return goober.OkResult(nil)
}
