package controllers

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	OK = 0
	Error = 1
	TokenExpired = 100
)

const (
	Success = "Success!"
	Fail = "Fail!"
)

type Response struct {
	Error ErrorType
	Data  interface{}
}

type ErrorType struct {
	Code    int
	Message string
	// UserMessage string
}

func (r *Response) SetErrorMessage(code int, msg string, userMsg string) {
	r.Error.Code = code
	r.Error.Message = msg
	// r.Error.UserMessage = userMsg
}

func (r *Response) SetData(d interface{}) {
	r.Data = d
}

func SetRsp(code int, d interface{}) (rsp *Response) {
	rsp = &Response{}
	var msg [3]interface{}
	var ok bool
	if msg, ok = msgMapping[code]; !ok {
		code = OK_INSERT_FAILED
	}
	msg, _ = msgMapping[code]
	rsp.SetErrorMessage(code, msg[1].(string), "")
	rsp.SetData(d)
	return rsp
}

func SetRspMsg(code int, msg string, d interface{}) (rsp *Response) {
	rsp = &Response{}
	rsp.SetErrorMessage(code, msg, "")
	rsp.SetData(d)
	return rsp
}

func Rsp(code int, data interface{}, a ...interface{}) (httpCode int, rsp *Response) {
	rsp = &Response{}
	var pigError [3]interface{}
	var ok bool
	if pigError, ok = msgMapping[code]; !ok {
		code = OK_INSERT_FAILED
		pigError, _ = msgMapping[code]
	}
	httpCode, _ = pigError[0].(int)
	paramsCnt, _ := pigError[2].(int)
	sprintfArgs := make([]reflect.Value, paramsCnt + 1)
	sprintfArgs[0] = reflect.ValueOf(pigError[1])
	for i := 1; i < paramsCnt + 1; i++ {
		sprintfArgs[i] = reflect.ValueOf(a[i - 1])
	}
	val := reflect.ValueOf(fmt.Sprintf).Call(sprintfArgs)[0]
	rsp.SetErrorMessage(code, val.String(), "")
	rsp.SetData(data)
	return
}

func MysqlError(err error, custom string) string {
	for key, value := range mysqlErrorMapping {
		if strings.Contains(err.Error(), key) {
			return custom + value
		}
	}
	return err.Error()
}

const (
	OK_INSERT_SUCCESS = 0
	OK_INSERT_FAILED = 1
	OK_TOKEN_FAILED = 2
)

var mysqlErrorMapping = map[string]string{
	"1062": "创建重复",
}

var msgMapping = map[int][3]interface{}{
	OK_INSERT_SUCCESS:              {"0", "请求成功", "0"},
	OK_INSERT_FAILED:               {"1", "请求失败", "0"},
	OK_TOKEN_FAILED:                {"2", "token错误", "0"},
}
