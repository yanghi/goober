package error

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
func GetMsg(c int) string {
	msg, k := em[c]
	if !k {
		return em[ErrUnknown]
	}

	return msg
}
