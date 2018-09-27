package action

import (
	"fmt"
	"github.com/labstack/echo"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"apidocserver/models"
	"apidocserver/xrom_mysql"
)

var CH_NUM = 1000

var (
	upgrader1 = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
var imMap = make(map[string]*websocket.Conn)
var channel = make(chan models.Msg ,CH_NUM)
var mutex sync.Mutex

func IMSever(c echo.Context) error {
	ws, err := upgrader1.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error("连接错误：",err)
		return err
	}
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()
		if err!=nil{
			c.Logger().Error("读取消息错误：",err)
			return err
		}
		fmt.Println(string(msg))
		go analysis(c,msg,ws)
	}
	return nil
}

func analysis(c echo.Context,msg []byte,ws *websocket.Conn) error {
	var imMsg models.ImMsg
	err := json.Unmarshal(msg,&imMsg)
	if err != nil{
		c.Logger().Error("json解析错误：",err)
		return err
	}
	if imMsg.MsgFromId == "" || imMsg.MsgType == "" {
		c.Logger().Error("msgFromId为空或者msgtype为空")
		m := "msgFromId为空或者msgtype为空"
		MsgHandle(c,imMsg.MsgFromId,imMsg.MsgId,[]byte(m))
		return nil
	}
	if imMsg.MsgType == "client" {
		imMap[imMsg.MsgFromId] = ws
		MsgHandle(c,imMsg.MsgFromId,imMsg.MsgId,msg)
		//fmt.Println("连接成功",imMsg.MsgFromId)
	}else if imMsg.MsgType == "p2p" {
		c.Logger().Error("消息：",imMsg)
		//把接收的数据插入表中
		xrom_mysql.InsertXORMMsg(imMsg)
		//把数据发送出去
		MsgHandle(c,imMsg.MsgFromId,imMsg.MsgId,msg)
		if ok := imMap[imMsg.MsgToId] != nil;ok {
			MsgHandle(c,imMsg.MsgToId,imMsg.MsgId,msg)
		}
	}else if imMsg.MsgType == "group" {
		groups := xrom_mysql.FindXROMMsg("SELECT * FROM im_group_member WHERE group_id = " + imMsg.MsgToId)
		if groups != nil {
			for _,group := range groups {
				//fmt.Println("待发送用户为：",string(group["user_id"]))
				MsgHandle(c,string(group["user_id"]),imMsg.MsgId,msg)
			}
		}
	} else if imMsg.MsgType == "logout" {

	}
	return nil
}

func MsgHandle(c echo.Context,toId string,msgId string,msg []byte)  {
	var newMsg models.Msg
	newMsg.Id = toId
	newMsg.MsgId = msgId
	newMsg.Msg = msg
	channel <- newMsg
	Broadcast(c,<-channel)
}

//发送信息
func Broadcast(c echo.Context,m models.Msg) error {
	if ok := imMap[m.Id] != nil;ok {
		mutex.Lock()
		err := imMap[m.Id].WriteMessage(websocket.TextMessage,m.Msg)
		if err != nil {
			mutex.Unlock()
			c.Logger().Error("发送信息错误")
			//fmt.Println("发送信息错误")
			delete(imMap,m.Id)
			return err
		}
		c.Logger().Debug("发送成功")
		//fmt.Println("发送成功")
		mutex.Unlock()
	}else {
		//保存离线消息,从redis取出未发送的消息，然后在重新填入
		/*var msgIds string
		ids := redispool.RedisGET("msg_"+m.Id)
		fmt.Println("离线消息保存ids",ids)
		if ids == nil{
			msgIds = m.MsgId
		}else{
			msgIds = string(ids) + "," + m.MsgId
		}
		redispool.RedisSETString("msg_"+m.Id,msgIds,0)*/
	}
	return nil
}
