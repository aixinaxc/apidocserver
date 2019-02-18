package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
	"crypto/md5"
	"encoding/hex"
)

func UserList(c echo.Context) error {
	pageNum := c.FormValue("page_num")
	pageSize := c.FormValue("page_size")
	pS,of := base.Offer(pageNum,pageSize)
	users,count := dao.UserList(pS,of)
	if users == nil {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,count,users))
}

func UserSave(c echo.Context) error {
	userId := c.FormValue("userId")
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	password = password + base.MD5
	pass := md5.New()
	pass.Write([]byte(password)) // 需要加密的字符串为buf.String()
	r := dao.UserSave(userId,username,hex.EncodeToString(pass.Sum(nil)))
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,r))
}


//删除API
func UserDelete(c echo.Context) error {
	userId := c.FormValue("userId")
	r := dao.UserDelete(userId)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}

//用户和项目绑定
func UserProjectSave(c echo.Context) error {
	userId := c.FormValue("userId")
	projectIds := c.FormValue("project_ids")
	if projectIds == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	r := dao.UserProjectSave(userId,projectIds)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}