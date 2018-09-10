package dao

import (
	"fmt"
	"apidocserver/xrom_mysql"
	"apidocserver/models"
	"apidocserver/base"
	"time"
)

//api列表
func ApiList(projectId string) []models.ApidocApi {
	engine := xrom_mysql.Client()
	apis := make([]models.ApidocApi,0)
	err:= engine.Cols("api_id", "sort_id","api_name").Asc("created_at").Where("project_id  = ? ",projectId).Find(&apis)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return apis
}

//api保存
func ApiSvae(apiId string,sortId string,projectId string,apiName string,apiEditContent string,apiShowContent string) string {
	engine := xrom_mysql.Client()
	api := new(models.ApidocApi)
	api.SortId = sortId
	api.ProjectId = projectId
	api.ApiName = apiName
	api.ApiEditContent = apiEditContent
	api.ApiShowContent = apiShowContent
	if apiId == "" {
		api.ApiId = base.UniqueId()
		api.ApiState = 1
		api.CreatedAt = int(time.Now().Unix())
		_,err := engine.Insert(&api)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return api.ApiId
	}else {
		api.UpdatedAt = int(time.Now().Unix())
		_, err := engine.Id(apiId).Update(api)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return string(projectId)
	}

}

func ApiContent(apiId string) models.ApidocApi {
	engine := xrom_mysql.Client()
	api := new(models.ApidocApi)
	b,err := engine.Where("api_id = ? ",apiId).Get(api)
	fmt.Println(b,err)
	return *api
}

func ApiDelete(apiId string) string {
	engine := xrom_mysql.Client()
	api := new(models.ApidocApi)
	b,err := engine.Id(apiId).Delete(api)
	fmt.Println(b,err)
	if err != nil {
		return "error"
	}
	return ""
}