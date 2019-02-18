package middle

import (
	"github.com/labstack/echo"
	"apidocserver/base"
	"apidocserver/redispool"
)

func LoginMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.FormValue("user_id")
		token := c.FormValue("token")
		if userId == "" || token == "" {
			return c.JSON(200,base.RetunMsgFunc(base.CodeUserNotLogin,0,nil))
		}
		redisToken := redispool.RedisGET(userId + "_ticket")
		if redisToken == nil || string(redisToken) == "" || string(redisToken) != token {
			return c.JSON(200,base.RetunMsgFunc(base.CodeUserNotLogin,0,nil))
		}
		return next(c)
	}
}