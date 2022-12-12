package msg

import "goblog/error/code"

var em = make(map[int]string)

func init() {
	em[code.InternalDefault] = "服务器错误"
	em[code.ParamsError] = "参数错误"
	em[code.Ok] = "ok"
	em[code.DB] = "数据库错误"
	em[code.Unknown] = "未知错误"
}
func GetMsg(c int) string {
	msg, k := em[c]
	if !k {
		return em[code.Unknown]
	}

	return msg
}
