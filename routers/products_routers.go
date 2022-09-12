package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/products"
	r "project/online-store/repository/products"
	s "project/online-store/service/products"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProductRoutes(echo *echo.Echo, conf config.Config) {
	db := config.IntiDB()

	repo := r.NewProductRepository(db)
	service := s.NewProductService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	g := echo.Group("/api")
	e := echo.Group("/api")

	// Middlewares
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(conf.JWT_SECRET),
		SigningMethod: "HS256",
	}))

	e.GET("/products", controll.FindAllProducts)
	e.GET("/products/:id", controll.FindProductByID)
	g.POST("/products", controll.CreateProduct)
	g.PUT("/products/:id", controll.UpdateProduct)
	g.DELETE("/products/:id", controll.DeleteProduct)
}
