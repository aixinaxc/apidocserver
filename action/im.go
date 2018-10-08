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
	total := c.FormValue("total")
	rm := new(base.ReturnMsg)
	if userFromId == "" || userToId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	fmt.Println("userFromId:",userFromId)
	fmt.Println("userToId:",userToId)
	fmt.Println("startTime:",startTime)
	fmt.Println("endTime:",endTime)
	msg,err := dao.MsgList(userFromId,userToId,startTime,endTime,msgType,total)
	if err != nil {
		fmt.Println("数据查询错误：",err)
		rm.Code401()
	}else {
		rm.Code200(0,msg)
	}
	return c.JSON(200,rm)
}
