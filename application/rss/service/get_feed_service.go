package service

import "goober/goober"

type GetFeedService struct {
	Href string `json:"href"`
}

func (s *GetFeedService) Get() *goober.ResponseResult {

	return goober.OkResult(nil)
}
