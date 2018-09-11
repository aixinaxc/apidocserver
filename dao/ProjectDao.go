package dao

import (
	"apidocserver/models"
	"apidocserver/xrom_mysql"
	"fmt"
	"apidocserver/base"
	"time"
)
//获得项目列表
func ProjectList(userId string) []models.ApidocProject {
	engine := xrom_mysql.Client()
	projects := make([]models.ApidocProject,0)
	var err error
	if userId == "" {
		err = engine.Asc("created_at").Find(&projects)
	}else {
		sql := "SELECT p.* FROM apidoc_project AS p,apidoc_user_project AS up WHERE p.project_id = up.project_id AND up.user_id = ? ORDER BY created_at ASC"
		err = engine.SQL(sql,userId).Find(&projects)
	}

	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return projects
}

//编辑项目
func ProjectSave(userId string,projectId string,projectName string) string {
	engine := xrom_mysql.Client()
	project := new(models.ApidocProject)
	if projectId == "" {
		project.ProjectId = base.UniqueId()
		project.ProjectName = projectName
		project.ProjectState = 1
		project.CreatedAt = int(time.Now().Unix())

		userProject := new(models.ApidocUserProject)
		userProject.UserId = userId
		userProject.ProjectId = project.ProjectId
		session := engine.NewSession()
		defer session.Close()
		err := session.Begin()
		_, err = session.Insert(project)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return "error"
		}
		_, err = session.Insert(userProject)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return "error"
		}
		err = session.Commit()
		if err != nil {
			fmt.Println(err)
			return "error"
		}
		return project.ProjectId
	}else {
		project.ProjectName = projectName
		project.UpdatedAt = int(time.Now().Unix())
		_, err := engine.Id(projectId).Update(project)
		if err != nil {
			fmt.Println("project_save:",err)
			return "error"
		}
		return projectId
	}
}


func ProjectDelete(projectId string) string {
	engine := xrom_mysql.Client()
	project := new(models.ApidocProject)
	b,err := engine.Id(projectId).Delete(project)
	fmt.Println(b,err)
	if err != nil {
		return "error"
	}
	return ""
}