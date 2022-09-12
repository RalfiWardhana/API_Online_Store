package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/auth"
	r "project/online-store/repository/auth"
	s "project/online-store/service/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthRoutes(echo *echo.Echo, conf config.Config) {
	db := config.IntiDB()

	repo := r.NewAuthRepository(db)
	service := s.NewAuthService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	e := echo.Group("/api")
	g := echo.Group("/api")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(conf.JWT_SECRET),
		SigningMethod: "HS256",
	}))

	e.POST("/register", controll.Register)
	e.POST("/login", controll.Login)

	g.GET("/store/activate", controll.ActivatedStore)
}
