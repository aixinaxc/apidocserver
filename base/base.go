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


type ReturnMsg struct {
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
}
