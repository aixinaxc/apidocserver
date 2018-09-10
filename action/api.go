package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
)


func ApiSvae(c echo.Context) error {
	apiId := c.FormValue("api_id")
	projectId := c.FormValue("project_id")
	sortId := c.FormValue("sortId")
	apiName := c.FormValue("api_name")
	apiEditContent := c.FormValue("api_edit_content")
	apiShowContent := c.FormValue("api_show_content")
	r := dao.ApiSvae(apiId,sortId,projectId,apiName,apiEditContent,apiShowContent)
	rm := new(base.ReturnMsg)
	if r == "error" {
		rm.Code401()
	}else {
		rm.Code200(1,r)
	}
	return c.JSON(200,rm)
}