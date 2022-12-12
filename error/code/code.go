// 通用错误码
//
// 前端错误码201-300
//
// 后端错误码301+
package code

const (
	// 内部错误码
	InternalDefault = 0
	DB              = -1
	Unknown         = -2

	Ok = 200
	// 前端错误码

	ParamsError = 201
	NotExsit    = 202
	UnExpect    = 203
	// 后端错误码

	ServerError = 301

	// 其他

	AuthExpired = 401
)
