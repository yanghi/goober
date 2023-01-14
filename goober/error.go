package goober

import "fmt"

const (
	// 内部错误码
	ErrInternalDefault = 0
	ErrDB              = -1
	ErrUnknown         = -2

	ErrOk = 200
	// 前端错误码

	ErrParamsInvlid = 201
	ErrNotExsit     = 202
	ErrUnExpect     = 203
	// 后端错误码

	ErrServerError = 301

	// 其他

	ErrTokenExpired = 401
	ErrTokenInvalid = 402
	ErrTokenMissing = 403
)

type Error struct {
	raw  error
	msg  string
	code int
}

func (e *Error) Error() string {
	if e.raw != nil {
		return e.raw.Error()
	}
	return e.msg
}

func (e *Error) Raw(err error) *Error {
	e.raw = err
	e.msg = err.Error()
	e.code = ErrUnExpect

	return e
}
func (e *Error) Code(c int) *Error {
	e.code = c
	return e
}
func (e *Error) Msg(m string) *Error {
	e.msg = m
	return e
}
func (e *Error) GetMsg() string {
	return e.msg
}
func (e *Error) GetCode() int {
	return e.code
}
func (e *Error) Output() *Error {

	if e.raw != nil {
		fmt.Println("[goober error]:"+e.msg+" ", e.raw)
	} else {
		fmt.Println("[goober error]: ", e.msg)
	}

	return e
}
func NewError() *Error {
	return &Error{}
}
func NewWithCode(code int) *Error {

	return &Error{
		code: code,
		msg:  getErrMsgWithCode(code),
	}
}

var em = make(map[int]string)

func init() {
	em[ErrInternalDefault] = "服务器错误"
	em[ErrParamsInvlid] = "参数错误"
	em[ErrOk] = "ok"
	em[ErrDB] = "数据库错误"
	em[ErrUnknown] = "未知错误"
	em[ErrTokenExpired] = "token已过期"
	em[ErrTokenInvalid] = "无效token"
	em[ErrTokenMissing] = "缺少token"
}

func toGError(e any) Error {
	var ge Error

	switch e.(type) {
	case string:
	case *Error:
		ge = *e.(*Error)
	case error:
		ge = Error{raw: e.(error)}
	}

	return ge
}

func getErrMsgWithCode(c int) string {
	msg, k := em[c]
	if !k {
		return em[ErrUnknown]
	}

	return msg
}
