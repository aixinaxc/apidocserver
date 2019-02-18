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
	projects,total := dao.ProjectList(userId)
	if projects == nil || len(projects) == 0 {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataEmpty,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,total,projects))
}

// 编辑项目
func ProjectSave(c echo.Context) error {
	userId := c.FormValue("user_id")
	projectId := c.FormValue("project_id")
	projectName := c.FormValue("project_name")
	fmt.Println("projectsave:",userId,projectId,projectName)
	if userId == "" ||  projectName == ""{
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	pId := dao.ProjectSave(userId,projectId,projectName)
	if pId == "error"{
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,pId))
}

//删除项目
func ProjectDelete(c echo.Context) error {
	projectId := c.FormValue("project_id")
	fmt.Println("peojectId",projectId)
	r := dao.ProjectDelete(projectId)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}