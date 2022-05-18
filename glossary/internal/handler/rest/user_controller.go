package rest

import (
	"fmt"
	"net/http"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/domain"
	j "secondhand_glossary/internal/middleware/jwt"
	"secondhand_glossary/internal/model"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	conf config.Config
	s    domain.UserService
}

// GetUsers godoc
// @Summary Get all users
// @Tags user
// @Produce json
// @Success 200 {object} []model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user [get]
func (c *UserController) GetUsersController(ctx echo.Context) error {
	users, err := c.s.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"users":   users,
	})
}

// DeleteUser godoc
// @Summary Delete user by id
// @Tags user
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/:user_id [delete]
func (c *UserController) DeleteUserController(ctx echo.Context) error {
	user_str := ctx.Param("user_id")
	user_id, err := strconv.Atoi(user_str)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}
	if err := c.s.DeleteUser(uint(user_id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting user",
	})
}

// Register godoc
// @Summary Register user
// @Tags user
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/register [post]
func (c *UserController) RegisterController(ctx echo.Context) error {
	data := model.UserRegister{}
	ctx.Bind(&data)

	user, err := c.s.Register(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	tokenDetails, err := j.CreateToken(user.ID, false, c.conf)
  if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
  }

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  tokenDetails.AccessToken,
		"refresh_token": tokenDetails.RefreshToken,
	})
}

// Login godoc
// @Summary Login user
// @Tags user
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/login [post]
func (c *UserController) LoginController(ctx echo.Context) error {
  // TODO Detect if token exist
	data := model.UserLogin{}
	ctx.Bind(&data)

	user, err := c.s.Login(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	tokenDetails, err := j.CreateToken(user.ID, false, c.conf)
  if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
  }

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  tokenDetails.AccessToken,
		"refresh_token": tokenDetails.RefreshToken,
	})
}

// func (c *UserController) LogoutController(ctx echo.Context) error {} // TODO Do this after redis
func (c *UserController) UpdateProfileController(ctx echo.Context) error {
	data := model.User{}

	ctx.Bind(&data)

	user, err := c.s.UpdateProfile(data)
  if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
    "user": user,
	})
}

// GetProfileDetails godoc
// @Summary Get user profiles
// @Tags user
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/profile [get]
func (c *UserController) GetProfileDetailsController(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*j.CustomClaims)
	userId := claims.UserID
  fmt.Println(userId)

	user, err := c.s.GetProfileDetails(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

