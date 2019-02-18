package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)

//获得api列表
func ApiList(c echo.Context) error {
	projectId := c.FormValue("project_id")
	if projectId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	api := dao.ApiContentList(projectId)
	if api == nil {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,api))
}


//获得api内容
func ApiContent(c echo.Context) error {
	apiId := c.FormValue("api_id")
	if apiId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	api := dao.ApiContent(apiId)
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,api))
}

//Api 保存
func ApiSvae(c echo.Context) error {
	apiId := c.FormValue("api_id")
	projectId := c.FormValue("project_id")
	sortId := c.FormValue("sort_id")
	apiName := c.FormValue("api_name")
	apiEditContent := c.FormValue("api_edit_content")
	apiShowContent := c.FormValue("api_show_content")
	if projectId == "" || sortId == "" || apiName == "" || apiEditContent == "" || apiShowContent == ""{
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	r := dao.ApiSvae(apiId,sortId,projectId,apiName,apiEditContent,apiShowContent)

	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,r))
}

//删除API
func ApiDelete(c echo.Context) error {
	apiId := c.FormValue("api_id")
	if apiId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	r := dao.ApiDelete(apiId)
	if r == "error" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataError,0,nil))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}