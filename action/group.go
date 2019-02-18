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
	if groupList == nil {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,groupList))
}
