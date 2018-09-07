package util

import "net/http"

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	OK     = "OK"
	FAILED = "FAILED"
)

func Ok(msg string, data interface{}) (httpStatus int, r Result) {
	httpStatus = http.StatusOK
	r.Msg = msg
	r.Data = data
	r.Code = OK
	return
}

func Fail(msg string, data interface{}) (httpStatus int, r Result) {
	httpStatus = http.StatusOK
	r.Msg = msg
	r.Data = data
	r.Code = FAILED
	return
}
