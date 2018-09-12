package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"apidocserver/action"
	"apidocserver/base"
	"apidocserver/middle"
)

func main() {
	e := echo.New()
	base.Config()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Static("/", "")
	e.POST("/login", action.Login)
	e.GET("/logout", action.Logout)

	e.GET("/project/list", action.ProjectList,middle.LoginMiddle)
	e.POST("/project/save", action.ProjectSave,middle.LoginMiddle)

	e.GET("/sort/list", action.SortList,middle.LoginMiddle)
	e.GET("/sort/sort_api_list", action.SortApiList,middle.LoginMiddle)
	e.POST("/sort/save", action.SortSave,middle.LoginMiddle)

	e.GET("/api/content", action.ApiContent,middle.LoginMiddle)
	e.POST("/api/save", action.ApiSvae,middle.LoginMiddle)
	e.GET("/api/delete", action.ApiDelete,middle.LoginMiddle)

	e.GET("/user/list", action.UserList,middle.LoginMiddle)
	e.POST("/user/save", action.UserSave,middle.LoginMiddle)
	e.GET("/user/delete", action.UserDelete,middle.LoginMiddle)
	e.POST("/user/user_project_save", action.UserProjectSave,middle.LoginMiddle)
	e.Logger.Fatal(e.Start(":9001"))
}

