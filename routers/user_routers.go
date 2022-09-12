package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/user"
	r "project/online-store/repository/user"
	s "project/online-store/service/user"

	"github.com/labstack/echo/v4"
)

func UserRoutes(echo *echo.Echo, conf config.Config) {
	db := config.IntiDB()

	repo := r.NewUserRepository(db)
	service := s.NewUserService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	g := echo.Group("/api")

	g.GET("/users", controll.FindAllUsers)
	g.GET("/users/:id", controll.FindUserById)

	g.PUT("/users/:id", controll.UpdateUserById)

	g.DELETE("/users/:id", controll.DeleteUserById)
}
