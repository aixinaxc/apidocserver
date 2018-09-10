package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)

//分类列表
func SortList(c echo.Context) error{
	projectId := c.FormValue("project_id")
	rm := new(base.ReturnMsg)
	if projectId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	sort := dao.SortList(projectId)
	rm.Code200(1,sort)
	return c.JSON(200,rm)
}

//侧边栏列表
func SortApiList(c echo.Context) error{
	projectId := c.FormValue("project_id")
	rm := new(base.ReturnMsg)
	if projectId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	sort := dao.SortList(projectId)
	api  := dao.ApiList(projectId)
	m := make(map[string]interface{},0)
	m["sort"] = sort
	m["api"] = api
	rm.Code200(1,m)
	return c.JSON(200,rm)
}

//分类保存
func SortSave(c echo.Context) error{
	sortId := c.FormValue("sort_id")
	projectId := c.FormValue("project_id")
	sortName := c.FormValue("sort_name")
	r := dao.SortSave(sortId,projectId,sortName)
	rm := new(base.ReturnMsg)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}