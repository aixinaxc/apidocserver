package base

import (
	"time"
	"crypto/rand"
	"encoding/base64"
	"io"
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"strconv"
	"fmt"
)

const MD5  = "12qw12er34sd"

const (
	MysqlMaxLifetime = 10*60*1000
	MysqlMaxOpenConn = 50
	MysqlMaxIdleConn = 1000
)

const (
	RedisMaxIdle = 3
	RedisMaxActive = 5
	RedisMIdleTimeout = 240 * time.Second
)

//分页
func Offer(pageNum,pageSize string) (int,int) {
	pN,err:=strconv.Atoi(pageNum)
	if err!= nil{
		fmt.Println("offer_err:",err)
		return 0,0
	}
	pS,err:=strconv.Atoi(pageSize)
	if err!= nil{
		fmt.Println("offer_err:",err)
		return 0,0
	}
	of := (pN-1) * pS

	return pS,of
}


//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//结构体转map
func Struct2Map(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		data[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	err = nil
	return
}


/*type ReturnMsg struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Data interface{} `json:"data"`
}

//返回200数据
func (rm *ReturnMsg) Code200(t int64,d interface{})  {
	rm.Code = 200
	rm.Msg = "OK"
	rm.Total = t
	rm.Data = d
}


//返回517数据
func (rm *ReturnMsg) Code517()  {
	rm.Code = 517
	rm.Msg = "empty"
	rm.Total = 0
	rm.Data = nil
}

//返回400数据
func (rm *ReturnMsg) Code400()  {
	rm.Code = 400
	rm.Msg = "data loss"
	rm.Total = 0
	rm.Data = nil
}

//返回401数据
func (rm *ReturnMsg) Code401()  {
	rm.Code = 401
	rm.Msg = "error"
	rm.Total = 0
	rm.Data = nil
}

//返回402数据
func (rm *ReturnMsg) Code402()  {
	rm.Code = 402
	rm.Msg = "user not logged in"
	rm.Total = 0
	rm.Data = nil
}*/


//--------Code返回码结构体
type Code struct {
	Code int
	Msg string
}

//--------定义返回码，以及返回信息
var (
	CodeDataSuccess = Code{Code:200,Msg:"success"}
	CodeDataEmpty = Code{Code:517,Msg:"empty"}
	CodeDataLoss = Code{Code:400,Msg:"data loss"}
	CodeDataError = Code{Code:401,Msg:"data error"}
	CodeUserNotLogin = Code{Code:402,Msg:"user not logged in"}
	CodeUserHas = Code{Code:403,Msg:"user has"}
)



//--------对需要返回的信息进行封装，方便对数据进行进一步处理
type ReturnMsg struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Total int64 `json:"total"`
	Data interface{} `json:"data"`
}

//--------对需要返回的信息进行赋值，并以结构体返回
func RetunMsgFunc(code Code,total int64,data interface{}) *ReturnMsg  {
	rm := new(ReturnMsg)
	rm.Code = code.Code
	rm.Msg = code.Msg
	rm.Total = total
	rm.Data = data
	return rm
}