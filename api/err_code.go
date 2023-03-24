package api

type ErrCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"error"`
}

func (ecode ErrCode) Error() string {
	return ecode.Msg
}

func (ecode ErrCode) WithErr(e error) ErrCode {
	ecode.Err = e
	return ecode
}

const (
	_RESERVE = iota + 999

	_CodeOk
	_CodeInternalErr
	_CodeDBErr
	_CodeParamErr
)

var (
	CodeOk                 = ErrCode{Code: _CodeOk, Msg: "请求成功"}
	CodeInternalServiceErr = ErrCode{Code: _CodeInternalErr, Msg: "服务器业务内部错误"}
	CodeDBErr              = ErrCode{Code: _CodeDBErr, Msg: "数据库错误"}
	CodeParamErr           = ErrCode{Code: _CodeParamErr, Msg: "请求参数有问题"}
)
