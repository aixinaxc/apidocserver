package dao

import (
	"fmt"
	"apidocserver/models"
	"apidocserver/xrom_mysql"
	"strconv"
)
//获得用户聊天列表
//p2p 点对点单聊
//group 群聊
func MsgList(userFromId string,userToId string,startTime string,endTime string,msgType string,total string) ([]models.ImMsg,error) {
	sql := " 1 = 1 "
	if userFromId != "" && userToId!=""{
		if msgType == "p2p" {
			sql = sql + " AND (msg_from_id = '" + userFromId +"' AND msg_to_id = '" + userToId + "') OR (msg_from_id = '" + userToId + "' AND msg_to_id = '"  + userFromId + "')"
		}else if msgType == "group" {
			sql = sql + " AND msg_to_id = '" + userToId + "'"
		}
	}
	if startTime != "0" {
		sql = sql + " AND created_at >= " + startTime
	}
	if endTime != "0" {
		sql = sql + " AND created_at < " + endTime
	}
	fmt.Println("sql:",sql)
	msg := make([]models.ImMsg, 0)
	engine := xrom_mysql.Client()
	if engine == nil {
		fmt.Println("mysql连接失败")
		return nil,nil
	}
	var err error
	if userFromId == "" && userToId == "" && startTime == "" && endTime == ""{
		err = engine.Asc("created_at").Find(&msg)
	}else {
		if  total == "0" {
			err = engine.Where(sql).Asc("created_at").Find(&msg)
		}else {
			itotal,err:=strconv.Atoi(total)
			if err!=nil {
				fmt.Println("转化错误：",err)
			}
			m := new(models.ImMsg)
			count, err := engine.Where(sql).Count(m)
			if err!=nil {
				fmt.Println("获取数量错误：",err)
			}
			err = engine.Where(sql).Asc("created_at").Limit(itotal,int(count)-itotal).Find(&msg)
		}

	}
	return msg,err
}


