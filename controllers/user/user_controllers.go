package user

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S user.UserService
}

// @Summary Find all users
// @Description Find all users
// @Router /api/users [get]
// @Success 200 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
func (c *Controller) FindAllUsers(ctx echo.Context) error {
	users, err := c.S.FindAllUsers()

	if len(users) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]any{
			"message": "No users found",
		})
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully found all users",
		"data":    users,
	})
}

// @Summary Find user by id
// @Description Find user by id
// @Router /api/users/{id} [get]
// @Param id path string true "User id"
// @Success 200 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
func (c *Controller) FindUserById(ctx echo.Context) error {
	// param id
	id := ctx.Param("id")
	// call service method to create user
	users, er := c.S.FindUserById(id)
	// if error throw error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}
	// return response
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully found user",
		"data":    users,
	})
}

// @Summary Update user by id
// @Description Update user by id
// @Router /api/users/{id} [put]
// @Param id path string true "User id"
// @Param user body entity.User true "User object"
// @Success 201 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
func (c *Controller) UpdateUserById(ctx echo.Context) error {
	// entity user
	user := entity.User{}
	// param id
	id := ctx.Param("id")
	// request body or bind user
	err := ctx.Bind(&user)
	// check if request body is empty throw error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}
	// call service method to update user
	er := c.S.UpdateUserById(id, user)
	// if error throw error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}
	// return response
	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": "Successfully updated user",
	})
}

// @Summary Delete user by id
// @Description Delete user by id
// @Router /api/users/{id} [delete]
// @Param id path string true "User id"
// @Success 200 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
func (c *Controller) DeleteUserById(ctx echo.Context) error {
	user := entity.User{}
	// param id
	id := ctx.Param("id")
	// call service method to delete user
	err := c.S.DeleteUserById(id, user)
	// if error throw error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   err.Error(),
			"message": "Internal server error",
		})
	}
	// return response
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully deleted user",
	})
}
