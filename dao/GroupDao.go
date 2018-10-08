package dao

import (
	"apidocserver/xrom_mysql"
	"apidocserver/models"
)

//获取用户群组
func GroupListByUserId(userId string) []models.ImGroup {
	engine := xrom_mysql.Client()
	groups := make([]models.ImGroup,0)
	engine.Join("INNER", "im_group_mumber", "im_group.group_id = im_group_mumber.group_id").Where(" im_group_mumber.user_id = ?",userId).Find(&groups)
	return groups
}