package dao

import (
	"apidocserver/xrom_mysql"
	"apidocserver/models"
	"fmt"
	"apidocserver/base"
	"time"
)

//获得用户列表
func UserList(pageSize int,offer int) ([]models.ApidocUser,int64){
	engine := xrom_mysql.Client()
	users := make([]models.ApidocUser,0)
	err:= engine.Cols("user_id", "user_username","created_at").Desc("created_at").Where("user_id  <> '1' ").Limit(pageSize,offer).Find(&users)
	user := new(models.ApidocUser)
	total, err := engine.Where("user_id  <> '1'").Count(user)
	if err!=nil {
		fmt.Println(err)
		return nil,0
	}
	return users,total
}

//根据用户名查找用户
func FindUserName(username string)  models.ApidocUser {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	b,err := engine.Where("user_username = ?",username).Get(user)
	fmt.Println(b,err)
	return *user
}


//查找指定用户
func FindUser(username string,password string) models.ApidocUser {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	b,err := engine.Cols("user_id","user_username").Where("user_username = ? AND user_password = ?",username,password).Get(user)
	fmt.Println(b,err)
	return *user
}

//保存用户
func UserSave(userId string,userName string,password string) string {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	user.UserUsername = userName
	user.UserPassword = password
	if userId == "" {
		user.UserId = base.UniqueId()
		user.UserState = 1
		user.CreatedAt = int(time.Now().Unix())
		_,err := engine.Insert(user)
		if err != nil {
			fmt.Println("sort_save:",err)
			return "error"
		}
		return user.UserId
	}else {
		user.UpdatedAt = int(time.Now().Unix())
		_, err := engine.Id(userId).Update(user)
		if err != nil {
			fmt.Println("user_save:",err)
			return "error"
		}
		return userId
	}
}

//删除用户
func UserDelete(userId string) string {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	b,err := engine.Id(userId).Delete(user)
	fmt.Println(b,err)
	if err != nil {
		return "error"
	}
	return ""
}