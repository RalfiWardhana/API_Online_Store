package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/category"
	r "project/online-store/repository/category"
	s "project/online-store/service/category"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CategoryRoutes(echo *echo.Echo, conf config.Config) {
	db := config.IntiDB()

	repo := r.NewCategoryRepository(db)
	service := s.NewCategoryService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	e := echo.Group("/api")
	g := echo.Group("/api")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(conf.JWT_SECRET),
		SigningMethod: "HS256",
	}))

	g.POST("/category", controll.CreateCategory)

	e.GET("/category", controll.FindAllCategory)
	e.GET("/category/:id", controll.FindCategoryByID)
}
