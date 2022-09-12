package main

import (
	"project/online-store/config"
	"project/online-store/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	conf := config.Config{}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	routers.AuthRoutes(e, conf)
	routers.UserRoutes(e, conf)
	routers.RoleRoutes(e, conf)
	routers.ProductRoutes(e, conf)
	routers.CategoryRoutes(e, conf)

	err := e.Start("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
