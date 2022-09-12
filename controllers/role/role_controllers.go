package role

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/role"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S role.RoleService
}

func (c *Controller) CreateRole(ctx echo.Context) error {
	// entity role
	role := entity.Role{}
	// request body or bind role
	err := ctx.Bind(&role)
	// check if request body is empty throw error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}
	// call service method to create role
	er := c.S.CreateRole(role)
	// if error throw error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}
	// return response
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully created role",
		"data":    role,
	})
}
