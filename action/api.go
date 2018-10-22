package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)

//获得api列表
func ApiList(c echo.Context) error {
	projectId := c.FormValue("project_id")
	rm := new(base.ReturnMsg)
	if projectId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	api := dao.ApiContentList(projectId)
	if api == nil {
		rm.Code401()
	}else {
		rm.Code200(0,api)
	}
	return c.JSON(200,rm)
}


//获得api内容
func ApiContent(c echo.Context) error {
	apiId := c.FormValue("api_id")
	rm := new(base.ReturnMsg)
	if apiId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	api := dao.ApiContent(apiId)
	rm.Code200(1,api)
	return c.JSON(200,rm)
}


func ApiSvae(c echo.Context) error {
	apiId := c.FormValue("api_id")
	projectId := c.FormValue("project_id")
	sortId := c.FormValue("sort_id")
	apiName := c.FormValue("api_name")
	apiEditContent := c.FormValue("api_edit_content")
	apiShowContent := c.FormValue("api_show_content")
	rm := new(base.ReturnMsg)
	if projectId == "" || sortId == "" || apiName == "" || apiEditContent == "" || apiShowContent == ""{
		rm.Code400()
		return c.JSON(200,rm)
	}
	r := dao.ApiSvae(apiId,sortId,projectId,apiName,apiEditContent,apiShowContent)

	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}

//删除API
func ApiDelete(c echo.Context) error {
	apiId := c.FormValue("api_id")
	rm := new(base.ReturnMsg)
	if apiId == "" {
		rm.Code400()
		return c.JSON(200,rm)
	}
	r := dao.ApiDelete(apiId)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(0,nil)
	}
	return c.JSON(200,rm)
}