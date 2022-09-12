package category

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S category.CategoryService
}

// @Summary Create category
// @Description Create category
// @Router /api/category POST
// @Success 200 {object} entity.Category
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
func (c *Controller) CreateCategory(ctx echo.Context) error {
	category := entity.Category{
		ID: uuid.New().String(),
	}

	err := ctx.Bind(&category)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	er := c.S.CreateCategory(category)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Category created successfully",
		"data":    category,
	})
}

// @Summary Find all categories
// @Description Find all categories
// @Router /api/category GET
// @Success 200 {object} entity.Category
// @Failure 500 {object} entity.Error
func (c *Controller) FindAllCategory(ctx echo.Context) error {
	category, err := c.S.FindAllCategory()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Categories found successfully",
		"data":    category,
	})
}

// @Summary Find category by ID
// @Description Find category by ID
// @Router /api/category/{id} GET
// @Param id path string true "Category ID"
// @Success 200 {object} entity.Category
// @Failure 404 {object} entity.Error
// @Failure 500 {object} entity.Error
func (c *Controller) FindCategoryByID(ctx echo.Context) error {
	id := ctx.Param("id")

	category, err := c.S.FindCategoryByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]any{
			"error":   err.Error(),
			"message": "Category not found",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Category found successfully",
		"data":    category,
	})
}
