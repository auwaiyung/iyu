package errcode

import (
	"fmt"
)

type ErrCode struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *ErrCode) IsSuccess() bool {
	if e.Code == 0 {
		return true
	}
	return false
}

func (e *ErrCode) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", e.Code, e.Msg, e.Data)
}

func (e *ErrCode) GetErrCode() *ErrCode {
	return e
}

func (e *ErrCode) WithData(data interface{}) *ErrCode {
	e.Data = data
	return e
}

func (e *ErrCode) AddMsg(msg ...interface{}) *ErrCode {
	e.Msg += ":" + fmt.Sprint(msg...)
	return e
}

func NewErrcode(code int, data interface{}) *ErrCode {
	return &ErrCode{
		Code: code,
		Data: data,
		Msg:  errCodeMap[code],
	}
}

const (
	ERROR   = iota - 1 // -1
	SUCCESS            // 0
	FAILURE
)

const (
	EMPTY_DATA = iota + 1000
	EXISTS_DATA
	NO_EXISTS_DATA
	ILLEGAL_PARAM
)

const (
	QUERY_ERROR = iota + 10000
	UPDATE_ERROR
	INSERT_ERROR
	DELETE_ERROR
)

// 常用
var (
	Error        = NewErrcode(ERROR, nil)
	Success      = NewErrcode(SUCCESS, nil)
	Failure      = NewErrcode(FAILURE, nil)
	EmptyData    = NewErrcode(EMPTY_DATA, nil)
	ExsistData   = NewErrcode(EXISTS_DATA, nil)
	NoExsistData = NewErrcode(NO_EXISTS_DATA, nil)
	IllegalParam = NewErrcode(ILLEGAL_PARAM, nil)

	QueryError  = NewErrcode(QUERY_ERROR, nil)
	UpdateError = NewErrcode(UPDATE_ERROR, nil)
	InsertError = NewErrcode(INSERT_ERROR, nil)
	DeleteError = NewErrcode(DELETE_ERROR, nil)
)

var errCodeMap = map[int]string{
	ERROR:          "异常",
	SUCCESS:        "成功",
	FAILURE:        "失败",
	EMPTY_DATA:     "数据为空",
	EXISTS_DATA:    "数据已存在",
	NO_EXISTS_DATA: "数据不存在",
	ILLEGAL_PARAM:  "参数不合法",

	QUERY_ERROR:  "查询错误",
	UPDATE_ERROR: "更新错误",
	INSERT_ERROR: "添加错误",
	DELETE_ERROR: "删除错误",
}
