package products

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/products"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S products.ProductService
}

// @Summary Find all products
// @Description Find all products
// @Routers /api/products GET
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
func (c *Controller) FindAllProducts(ctx echo.Context) error {
	products, err := c.S.FindAllProducts()
	if len(products) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]any{
			"message": "No products found",
		})
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully found products",
		"data":    products,
	})
}

// @Summary Create product
// @Description Create product
// @Routers /api/products POST
// @Param product body entity.Product true "Product"
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
func (c *Controller) CreateProduct(ctx echo.Context) error {
	user := ctx.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	roleID := claims["role"].(float64)
	UserID := claims["id"].(string)

	if roleID == 1 {
		return ctx.JSON(http.StatusUnauthorized, map[string]any{
			"message": "You are not authorized to create products",
		})
	}

	product := entity.Product{
		Id:     uuid.New().String(),
		UserId: UserID,
	}

	err := ctx.Bind(&product)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	er := c.S.CreateProduct(product)

	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong",
			"error":   er.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": "Successfully created product",
		"data":    product,
	})
}

// @Summary Find product by ID
// @Description Find product by ID
// @Routers /api/products/{id} GET
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
func (c *Controller) FindProductByID(ctx echo.Context) error {
	id := ctx.Param("id")

	product, err := c.S.FindProductByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]any{
			"message": "Product not found",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully found product",
		"data":    product,
	})
}

// @Summary Update product
// @Description Update product
// @Routers /api/products/{id} PUT
// @Param id path string true "Product ID"
// @Param product body entity.Product true "Product"
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Security JwtAuth
func (c *Controller) UpdateProduct(ctx echo.Context) error {
	product := entity.Product{}

	id := ctx.Param("id")

	err := ctx.Bind(&product)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	products, er := c.S.UpdateProduct(id, product)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong",
			"error":   er.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully updated product",
		"data":    products,
	})
}

// @Summary Delete product
// @Description Delete product
// @Routers /api/products/{id} DELETE
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Security JwtAuth
func (c *Controller) DeleteProduct(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.S.DeleteProduct(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully deleted product",
	})

}
