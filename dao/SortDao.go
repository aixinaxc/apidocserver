package dao

import (
	"fmt"
	"apidocserver/xrom_mysql"
	"apidocserver/base"
	"time"
	"apidocserver/models"
)

//api分类列表
func SortList(projectId string) []models.ApidocSort {
	engine := xrom_mysql.Client()
	sorts := make([]models.ApidocSort,0)
	err:= engine.Cols("sort_id", "sort_name").Asc("created_at").Where("project_id = ?",projectId).Find(&sorts)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return sorts
}

//保存api分类
func SortSave(sortId string,projectId string,sortName string) string {
	engine := xrom_mysql.Client()
	sort := new(models.ApidocSort)
	if sortId == "" {
		sort.SortId = base.UniqueId()
		sort.ProjectId = projectId
		sort.SortName = sortName
		sort.Status = 1
		sort.CreatedAt = int(time.Now().Unix())
		_,err := engine.Insert(sort)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return sort.SortId
	}else {
		sort.SortName = sortName
		sort.UpdatedAt = int(time.Now().Unix())
		_, err := engine.Id(sortId).Update(sort)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return sortId
	}
}