package action

import (
	"github.com/labstack/echo"
	"apidocserver/dao"
	"apidocserver/base"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"apidocserver/redispool"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataLoss,0,nil))
	}
	password = password + base.MD5
	pass := md5.New()
	pass.Write([]byte(password)) // 需要加密的字符串为buf.String()
	user := dao.FindUser(username,hex.EncodeToString(pass.Sum(nil)))
	if user.UserUsername != "" {
		uuid := base.UniqueId()
		fmt.Println(uuid)
		userMap,_ := base.Struct2Map(user)
		userMap["Token"] = uuid
		redispool.RedisSETString(user.UserId + "_ticket",uuid,0)
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,1,userMap))
	}
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataEmpty,0,nil))
}

func Logout(c echo.Context) error {
	userId := c.FormValue("user_id")
	if userId == "" {
		return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
	}
	redispool.RedisDel(userId+"token")
	return c.JSON(200,base.RetunMsgFunc(base.CodeDataSuccess,0,nil))
}