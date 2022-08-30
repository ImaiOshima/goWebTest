package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success = NewError(0,"成功")
	ServerError = NewError(1000000,"服务内部错误")
	InvalidParam = NewError(1000001,"入参错误")
	NotFound = NewError(1000002,"找不到")
	UnauthorizedAuthNotExist = NewError(1000003,"鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError = NewError(1000004,"鉴权失败,Token失败")
	UnauthorizedTokenTimeOut= NewError(1000005,"鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(1000006,"鉴权失败,Token生成失败")
	TooManyRequest = NewError(1000007,"请求过多")
	)

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int,msg string) *Error{
	if _,ok := codes[code];ok{
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个",code))
	}
	codes[code] = msg
	return &Error{code:code,msg:msg}
}

func (e *Error) Error() string{
	return fmt.Sprintf("错误码:%d,错误信息:%s",e.code,e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string{
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string{
	return fmt.Sprintf(e.msg,args...)
}

func (e *Error) Details() []string{
	return e.details
}

func (e *Error) WithDetail(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _,d := range details {
		newError.details = append(newError.details,d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code(){
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParam.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeOut.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

