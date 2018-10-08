package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)

//获得群组列表
func GroupList(c echo.Context) error{
	userId := c.FormValue("user_id")
	groupList := dao.GroupListByUserId(userId)
	rm := new(base.ReturnMsg)
	if groupList == nil {
		rm.Code401()
	}else {
		rm.Code200(0,groupList)
	}
	return c.JSON(200,rm)
}
