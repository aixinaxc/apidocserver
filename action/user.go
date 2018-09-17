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
	rm := new(base.ReturnMsg)
	if users == nil {
		rm.Code401()
	}else {
		rm.Code200(count,users)
	}
	return c.JSON(200,rm)
}

func UserSave(c echo.Context) error {
	userId := c.FormValue("userId")
	username := c.FormValue("username")
	password := c.FormValue("password")
	rm := new(base.ReturnMsg)
	if username == "" || password == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	password = password + base.MD5
	pass := md5.New()
	pass.Write([]byte(password)) // 需要加密的字符串为buf.String()
	r := dao.UserSave(userId,username,hex.EncodeToString(pass.Sum(nil)))
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}


//删除API
func UserDelete(c echo.Context) error {
	userId := c.FormValue("userId")
	rm := new(base.ReturnMsg)
	r := dao.UserDelete(userId)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(0,nil)
	}
	return c.JSON(200,rm)
}


func UserProjectSave(c echo.Context) error {
	userId := c.FormValue("userId")
	projectIds := c.FormValue("project_ids")
	rm := new(base.ReturnMsg)
	if projectIds == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	r := dao.UserProjectSave(userId,projectIds)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(0,nil)
	}
	return  c.JSON(200,rm)
}