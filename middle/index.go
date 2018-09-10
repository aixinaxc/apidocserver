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
		rm := new(base.ReturnMsg)
		if userId == "" || token == "" {
			rm.Code402()
			return c.JSON(200,rm)
		}
		redisToken := redispool.RedisGET(userId + "_ticket")

		if redisToken == nil || string(redisToken) == "" || string(redisToken) != token {
			rm.Code402()
			return c.JSON(200,rm)
		}
		return next(c)
	}
}