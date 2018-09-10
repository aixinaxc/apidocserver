package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)

func UserList(c echo.Context) error {
	users := dao.UserList()
	userProjects := dao.UserProjectList("")
	rm := new(base.ReturnMsg)
	if users == nil || userProjects == nil {
		rm.Code401()
	}else {
		m := make(map[string]interface{},0)
		m["user"] = users
		m["user_project"] = userProjects
		rm.Code200(1,m)
	}
	return c.JSON(200,rm)
}

func UserSave(c echo.Context) error {
	userId := c.FormValue("user_id")
	username := c.FormValue("username")
	password := c.FormValue("password")
	r := dao.UserSave(userId,username,password)
	rm := new(base.ReturnMsg)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}


//删除API
func UserDelete(c echo.Context) error {
	userId := c.FormValue("user_id")
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
	userId := c.FormValue("user_id")
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