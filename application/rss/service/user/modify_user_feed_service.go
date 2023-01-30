package user

import (
	"goober/application/mysql"
	"goober/goober"

	"github.com/huandu/go-sqlbuilder"
)

type ModifyFeedService struct {
	UserId int
	ItemId int  `json:"itemId" form:"itemId"`
	Read   bool `json:"read" form:"read"`
}

func (s *ModifyFeedService) MarkReadState() *goober.ResponseResult {
	ub := sqlbuilder.NewUpdateBuilder()

	ub.Set("read=" + ub.Args.Add(s.Read))
	return s.modify(ub)
}

func (s *ModifyFeedService) modify(ub *sqlbuilder.UpdateBuilder) *goober.ResponseResult {
	ub.Update("gb_user_rss_feed_items").
		Where("uid="+ub.Args.Add(s.UserId), "item_id="+ub.Args.Add(s.ItemId))

	sql, args := ub.Build()

	_, e := mysql.DB().Exec(sql, args...)

	if e != nil {
		return goober.ErrorLogResponse(e, "修改用户item").Result()
	}

	return goober.OkResult(nil)
}
