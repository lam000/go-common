package ecode

var (
	OK = Add(0, "ok")

	ServerError        = Add(500, "系统错误，请稍后重试")
	RequestParaInvalid = Add(400, "请求参数错误")
)
