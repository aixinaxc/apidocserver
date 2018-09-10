package dao

import (
	"apidocserver/xrom_mysql"
	"apidocserver/models"
	"fmt"
)

func FindUser(username string,password string) models.ApidocUser {
	engine := xrom_mysql.Client()
	user := new(models.ApidocUser)
	b,err := engine.Cols("user_id","user_username").Where("user_username = ? AND user_password = ?",username,password).Get(user)
	fmt.Println(b,err)
	return *user
}