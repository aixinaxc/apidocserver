package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
	"strconv"
)

//分类分页列表
func SortPList(c echo.Context) error{
	projectId := c.FormValue("project_id")
	pageNum := c.FormValue("page_num")
	pageSize := c.FormValue("page_size")
	rm := new(base.ReturnMsg)
	if projectId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	pS,of := base.Offer(pageNum,pageSize)
	sorts,count := dao.SortPList(pS,of,projectId)
	rm.Code200(count,sorts)
	return c.JSON(200,rm)
}



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
	sortPid := c.FormValue("sort_pid")
	sortSeq := c.FormValue("sort_seq")
	iSortSeq,err := strconv.Atoi(sortSeq)
	rm := new(base.ReturnMsg)
	if err != nil || projectId == "" || sortName == ""{
		rm.Code400()
		return c.JSON(200,rm)
	}
	r := dao.SortSave(sortId,projectId,sortName,sortPid,iSortSeq)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}

//分类删除
func SortDelete(c echo.Context) error{
	sortId := c.FormValue("sort_id")
	rm := new(base.ReturnMsg)
	r := dao.SortDelete(sortId)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(0,nil)
	}
	return c.JSON(200,rm)
}