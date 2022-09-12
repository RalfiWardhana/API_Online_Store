package auth

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/auth"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S auth.AuthService
}

func (c *Controller) Register(ctx echo.Context) error {
	user := entity.User{
		Id: uuid.New().String(),
	}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	data, er := c.S.Register(user)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "User registered successfully",
		"data":    data,
	})
}

func (c *Controller) Login(ctx echo.Context) error {
	user := entity.User{
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	token, er := c.S.Login(user.Email, user.Password)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "User logged in successfully",
		"data":    token,
	})
}

func (c *Controller) ActivatedStore(ctx echo.Context) error {
	// update role user to admin
	user := ctx.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	err := c.S.ActivatedStore(userID, 2)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully activated store",
	})
}
