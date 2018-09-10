package dao

import (
	"apidocserver/xrom_mysql"
	"apidocserver/models"
	"fmt"
)

func FindUser(username string,password string) models.ApidocUser {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	b,err := engine.Where("user_username = ? AND user_password = ?",username,password).Get(user)
	fmt.Println(b,err)
	user.UserPassword = ""
	return *user
}