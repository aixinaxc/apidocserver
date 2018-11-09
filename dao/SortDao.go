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
	err:= engine.Cols("sort_id", "sort_name","sort_pid").Desc("sort_seq").Where("project_id = ?",projectId).Find(&sorts)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return sorts
}

//api分页列表
func SortPList(pageSize int,offer int,projectId string) ([]models.ApidocSort,int64) {
	engine := xrom_mysql.Client()
	sorts := make([]models.ApidocSort,0)
	err:= engine.Desc("created_at").Where("project_id = ?",projectId).Limit(pageSize,offer).Find(&sorts)
	sort := new(models.ApidocSort)
	total, err := engine.Where("project_id  = ?" ,projectId).Count(sort)
	if err!=nil {
		fmt.Println(err)
		return nil,0
	}
	return sorts,total
}


//保存api分类
func SortSave(sortId string,projectId string,sortName string,sortPid string,sortSeq int) string {
	engine := xrom_mysql.Client()
	sort := new(models.ApidocSort)
	if sortId == "" {
		sort.SortId = base.UniqueId()
		sort.ProjectId = projectId
		sort.SortName = sortName
		sort.SortPid = sortPid
		sort.SortSeq = sortSeq
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
		sort.SortPid = sortPid
		sort.SortSeq = sortSeq
		sort.UpdatedAt = int(time.Now().Unix())
		_, err := engine.Id(sortId).Update(sort)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return sortId
	}
}


//删除目录
func SortDelete(sortId string) string {
	engine := xrom_mysql.Client()
	sort := new(models.ApidocSort)
	b,err := engine.Id(sortId).Delete(sort)
	fmt.Println(b,err)
	if err != nil {
		return "error"
	}
	return ""
}