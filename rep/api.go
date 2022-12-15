package rep

import (
	gerr "goblog/error"
)

// api response
type Response struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}

func Build(data any, code int, msg string) *Response {

	r := Response{
		Data: data,
		Msg:  msg,
		Code: code,
		Ok:   code == gerr.ErrOk,
	}

	return &r
}

func BuildOkResponse(data any) *Response {
	return &Response{
		Data: data,
		Code: gerr.ErrOk,
		Msg:  "ok",
		Ok:   true,
	}
}

func BuildFatalResponse(e any) *Response {
	var ge gerr.GError

	switch e.(type) {
	case string:
	case error:
		ge = *gerr.New(e)
	case gerr.GError:
		ge = e.(gerr.GError)
	}

	return &Response{
		Msg:  ge.Msg(),
		Code: ge.Code,
	}
}
