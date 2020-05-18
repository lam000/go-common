package ecode

import (
	"fmt"
	"github.com/lam000/go-common/library/util/recovery"
	"github.com/pkg/errors"
)

var (
	_codes = map[int64]string{}
)

type ErrCode struct {
	code    int64
	message string
	detail  string
	stack   []byte
}

func New(e int64, msg string) ErrCode {
	if e < 0 {
		panic("business ecode must greater than zero")
	}

	return ErrCode{
		code:    e,
		message: msg,
	}
}

func Add(code int64, msg string) ErrCode {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", code))
	}

	_codes[code] = msg
	return ErrCode{
		code:    code,
		message: msg,
	}
}

func (ec ErrCode) Error() string {
	return ec.message
}

func (ec ErrCode) Code() int64 {
	return ec.code
}

func (ec ErrCode) Message() string {
	return ec.message
}

func (ec ErrCode) Stack() []byte {
	return ec.stack
}

func (ec ErrCode) Detail() string {
	return ec.detail
}

func (ec ErrCode) WithDetail(detail string) ErrCode {
	return ErrCode{
		code:    ec.Code(),
		message: ec.Message(),
		stack:   ec.Stack(),
		detail:  detail,
	}
}

func (ec ErrCode) WithStack() ErrCode {
	return ErrCode{
		code:    ec.Code(),
		message: ec.Message(),
		detail:  ec.Detail(),
		stack:   recovery.Stack(2),
	}
}

func Cause(e error) ErrCode {
	if e == nil {
		return OK
	}

	ec, ok := errors.Cause(e).(ErrCode)
	if ok {
		return ec
	}

	return ServerError
}

func Equal(a, b ErrCode) bool {
	return a.Code() == b.Code()
}

func EqualError(code ErrCode, err error) bool {
	return Cause(err).Code() == code.Code()
}
