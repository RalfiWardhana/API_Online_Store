package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/role"
	r "project/online-store/repository/role"
	s "project/online-store/service/role"

	"github.com/labstack/echo/v4"
)

func RoleRoutes(echo *echo.Echo, conf config.Config) {
	db := config.IntiDB()

	repo := r.NewRoleRepository(db)
	service := s.NewRoleService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	g := echo.Group("/api")

	g.POST("/roles", controll.CreateRole)
}
