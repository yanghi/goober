package service

import (
	"goober/goober"
	"time"

	"github.com/robfig/cron/v3"
)

type FeedUpdateJobService struct {
	TimmingSpec string
	entryId     cron.EntryID
	c           *cron.Cron
}

func (s *FeedUpdateJobService) Start() {

	if s.TimmingSpec == "" {
		s.TimmingSpec = "@every 2h"
		// s.TimmingSpec = "@every 1m"
	}

	c := cron.New()

	// id, e := c.AddFunc(s.TimmingSpec, func() { goober.Logger().Println("Every hour on the half hour") })
	id, e := c.AddFunc(s.TimmingSpec, s.run)

	if e != nil {
		goober.Logger().Errorln("[goober feed job] 创建定时任务失败", e)
		return
	}
	s.c = c
	s.entryId = id

	c.Start()
	goober.Logger().Println("[goober feed job] 准备就绪 ", s.TimmingSpec)

}
func (s *FeedUpdateJobService) Stop() {
	if s.c != nil {
		s.c.Remove(s.entryId)
	}
}
func (s *FeedUpdateJobService) run() {
	var usrv = UpdateFeedItemsService{}

	goober.Logger().Println("[goober feed job] 开始执行", time.Now().Format("2006-01-02T15:04:05"))

	r := usrv.UpdateAll()
	if r.Ok {
		data := r.Data.(UpdateFeedResult)

		goober.Logger().Println("[goober feed job] 执行成功", time.Now().Format("2006-01-02T15:04:05"))
		goober.Logger().Printf(
			"  已更新%d个订阅源,%d个成功,%d个失败,%d个订阅源有新数据,共更新%d条数据",
			data.TotalCount, data.SuccessCount, data.FailedCount, data.ValidCount, data.UpdateItemCount,
		)
	} else {
		goober.Logger().Errorln("[goober feed job] 执行失败", time.Now().Format("2006-01-02T15:04:05"))
		goober.Logger().Errorf("  code: %d, msg:%s\n", r.Code, r.Msg)
	}

}
