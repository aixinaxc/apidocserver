package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
	"fmt"
)

//项目列表
func ProjectList(c echo.Context) error {
	userId := c.FormValue("userId")
	projects := dao.ProjectList(userId)
	rm := new(base.ReturnMsg)
	if projects == nil || len(projects) == 0 {
		rm.Code517()
	}else {
		rm.Code200(len(projects),projects)
	}
	return c.JSON(200,rm)
}

// 编辑项目
func ProjectSave(c echo.Context) error {

	userId := c.FormValue("user_id")
	projectId := c.FormValue("project_id")
	projectName := c.FormValue("project_name")
	fmt.Println("projectsave:",userId,projectId,projectName)
	rm := new(base.ReturnMsg)
	if userId == "" || projectId == "" || projectName == ""{
		rm.Code400()
		return c.JSON(200,rm)
	}
	pId := dao.ProjectSave(userId,projectId,projectName)

	if pId == "error"{
		rm.Code401()
	}else {
		rm.Code200(1,pId)
	}
	return nil
}

//删除项目
func ProjectDelete(c echo.Context) error {

	return nil
}