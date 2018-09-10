package dao

import (
	"fmt"
	"apidocserver/xrom_mysql"
	"apidocserver/models"
)

//api列表
func ApiList(projectId string) []models.ApidocApi {
	engine := xrom_mysql.Client()
	apis := make([]models.ApidocApi,0)
	sql := "SELECT p.* FROM apidoc_project AS p,apidoc_user_project AS up WHERE p.project_id = up.project_id AND up.user_id = ?"
	err:= engine.SQL(sql,projectId).Find(&apis)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return apis
}
