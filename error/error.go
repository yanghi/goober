package error

// type GError error

type GError struct {
	raw  error
	msg  string
	Code int
}

func NewWithCode(code int) *GError {

	return &GError{
		Code: code,
		msg:  GetMsg(code),
	}
}

func New(args ...any) *GError {
	var code int
	var msg string
	var raw error

	for _, v := range args {
		switch v.(type) {
		case int:
			code = v.(int)
			// break
		case string:
			msg = v.(string)
		default:
			raw = v.(error)
		}
	}

	if msg == "" {
		msg = GetMsg(code)
	}

	return &GError{Code: code, msg: msg, raw: raw}
}

func Convert(e error) *GError {

	if ge, b := e.(*GError); b {
		return ge
	}

	return New(e)
}

func (e *GError) Error() string {
	if e.raw != nil {
		return e.raw.Error()
	}

	return e.msg
}
func (e *GError) Msg() string {
	if e.raw != nil {
		return e.raw.Error()
	}
	return e.msg
}
