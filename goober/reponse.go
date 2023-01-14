package goober

import "fmt"

type ResponseResult struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}

// api响应构造
type Response struct {
	res       ResponseResult
	err       error
	alredyLog bool
	log       bool
	label     string
}

func NewResponse() *Response {
	return &Response{}
}
func ErrorLogResponse(e error, label string) *Response {
	r := &Response{log: true, label: label, res: ResponseResult{}}

	r.RawError(e)
	return r
}

func (r *Response) From(res ResponseResult) *Response {
	r.res = res
	return r
}
func (r *Response) AnyError(e any) *Response {
	ge := toGError(e)

	r.Error(&ge)
	return r
}
func (r *Response) Error(err *Error) *Response {
	r.err = err
	r.res.Msg = err.Error()
	r.res.Code = err.GetCode()

	return r
}
func (r *Response) RawError(err error) *Response {
	r.err = err
	r.res.Msg = err.Error()
	r.res.Code = ErrUnExpect

	return r
}

func (r *Response) Ok() *Response {
	r.res.Ok = true
	return r
}
func (r *Response) Code(c int) *Response {
	r.res.Code = c
	return r
}
func (r *Response) Msg(msg string) *Response {
	r.res.Msg = msg
	return r
}
func (r *Response) Data(data any) *Response {
	r.res.Data = data
	return r
}
func (r *Response) OkResult() *ResponseResult {
	r.res.Ok = true
	if r.res.Msg == "" {
		r.res.Msg = "ok"
	}
	return r.Result()
}

func (r *Response) Result() *ResponseResult {
	if r.log {
		r.Log()
	}
	return &ResponseResult{
		Ok:   r.res.Ok,
		Code: r.res.Code,
		Msg:  r.res.Msg,
		Data: r.res.Data,
	}
}
func (r *Response) Label(label string) *Response {
	r.label = label
	return r
}
func (r *Response) AllowLog() *Response {
	r.log = true
	return r
}
func (r *Response) Log() *Response {

	if r.alredyLog || r.res.Ok {
		return r
	}

	r.alredyLog = true
	l := r.label
	if l != "" {
		l = "(" + l + ")"
	}
	fmt.Printf("[goobger response%s]  %s", l, r.res.Msg)

	if r.err != nil {
		fmt.Println(" error->", r.err)
	}

	return r
}
func (r *Response) FailedResult() *ResponseResult {
	if r.res.Msg == "" {
		if r.err != nil {
			r.res.Msg = r.err.Error()
		} else {
			r.res.Msg = "failed"
		}
	}

	return r.Result()
}

func OkResult(data any) *ResponseResult {
	return &ResponseResult{
		Ok:   true,
		Msg:  "ok",
		Code: 200,
		Data: data,
	}
}
func FailedResult(code int) *ResponseResult {
	var r = &ResponseResult{
		Ok:   false,
		Code: code,
	}

	r.Msg = getErrMsgWithCode(code)

	if r.Msg == "" {
		r.Msg = "failed"
	}
	return r
}
func WrongResult(err error) *ResponseResult {
	var r = Response{log: true}
	r.AnyError(err)
	return r.Result()
}
