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

func FatalResponseWithCode(c int) *Response {
	return &Response{
		Code: c,
		Msg:  gerr.GetMsg(c),
		Ok:   false,
	}
}
func FatalResponseWithGError(err gerr.GError) *Response {
	return &Response{
		Code: err.Code,
		Msg:  err.Msg(),
		Ok:   false,
	}
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
	case gerr.GError:
		ge = e.(gerr.GError)
	case error:
		ge = *gerr.New(e)

	}
	ge = *gerr.New(e)

	return &Response{
		Msg:  ge.Msg(),
		Code: ge.Code,
	}
}
