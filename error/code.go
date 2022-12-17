// 通用错误码
//
// 前端错误码201-300
//
// 后端错误码301+
package error

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
