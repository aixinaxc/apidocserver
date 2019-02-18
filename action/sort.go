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
	if projectId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	pS,of := base.Offer(pageNum,pageSize)
	sorts,count := dao.SortPList(pS,of,projectId)
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,count,sorts))
}

//分类列表
func SortList(c echo.Context) error{
	projectId := c.FormValue("project_id")
	if projectId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	sort := dao.SortList(projectId)
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,sort))
}

//侧边栏列表
func SortApiList(c echo.Context) error{
	projectId := c.FormValue("project_id")
	if projectId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	sort := dao.SortList(projectId)
	api  := dao.ApiList(projectId)
	m := make(map[string]interface{},0)
	m["sort"] = sort
	m["api"] = api
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,m))
}

//分类保存
func SortSave(c echo.Context) error{
	sortId := c.FormValue("sort_id")
	projectId := c.FormValue("project_id")
	sortName := c.FormValue("sort_name")
	sortPid := c.FormValue("sort_pid")
	sortSeq := c.FormValue("sort_seq")
	iSortSeq,err := strconv.Atoi(sortSeq)
	if err != nil || projectId == "" || sortName == ""{
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	r := dao.SortSave(sortId,projectId,sortName,sortPid,iSortSeq)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,r))
}

//分类删除
func SortDelete(c echo.Context) error{
	sortId := c.FormValue("sort_id")
	r := dao.SortDelete(sortId)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}