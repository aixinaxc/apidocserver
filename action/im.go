package action

import (
	"fmt"

	"github.com/labstack/echo"
	"apidocserver/base"
	"apidocserver/dao"
)

func MsgList(c echo.Context) error {
	userFromId := c.FormValue("msg_from_id")
	userToId := c.FormValue("msg_to_id")
	startTime := c.FormValue("start_time")
	endTime := c.FormValue("end_time")
	msgType := c.FormValue("msg_type")
	pageNum := c.FormValue("page_num")
	pageSize := c.FormValue("page_size")
	if userFromId == "" || userToId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	fmt.Println("userFromId:",userFromId)
	fmt.Println("userToId:",userToId)
	fmt.Println("startTime:",startTime)
	fmt.Println("endTime:",endTime)
	msg,count,err := dao.MsgList(userFromId,userToId,startTime,endTime,msgType,pageNum,pageSize)
	if err != nil {
		fmt.Println("数据查询错误：",err)
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,count,msg))
}
