package serror

import "github.com/pw1992/weixin/kernel/slog"

type Error struct {
	ErrMsg     string        `json:"err_msg"`
	ErrCode    int           `json:"err_code"`
	ErrContext []interface{} `json:"err_context"`
}

func NewError(msg string, code int, args ...interface{}) *Error {
	e := &Error{
		ErrMsg:     msg,
		ErrCode:    code,
		ErrContext: args,
	}
	e.log()
	return e
}

// 记录错误日志
func (e *Error) log() {
	log := slog.New()
	log.Error(e.ErrMsg, e.ErrContext)
}

// 主动抛出错误
func (e *Error) Throw() {
	panic(e)
}
