package dao

import (
	"fmt"
	"apidocserver/xrom_mysql"
	"strings"
	"apidocserver/models"
)

//获得用户项目关系
func UserProjectList(userId string) []models.ApidocUserProject {
	engine := xrom_mysql.Client()
	userProjects := make([]models.ApidocUserProject,0)
	var err error
	if userId == "" {
		err= engine.Find(&userProjects)
	}else {
		err= engine.Where("user_id = ?",userId).Find(&userProjects)
	}
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return userProjects
}

//保存用户项目关系
func UserProjectSave(userId string,projectIds string) string {
	projectIdArr := strings.Split(projectIds,",")
	engine := xrom_mysql.Client()
	userProject := new(models.ApidocUserProject)
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Where("user_id = ?",userId).Delete(userProject)
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return "error"
	}
	for _,up := range projectIdArr {
		userProject1 := new(models.ApidocUserProject)
		userProject1.UserId = userId
		userProject1.ProjectId = up
		_, err = session.Insert(userProject1)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return "error"
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	return ""
}
