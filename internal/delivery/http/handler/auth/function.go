package auth

import (
	"fmt"
	"net/http"

	md "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/delivery/http/middleware"
	eu "github.com/berrylradianh/cmlabs-backend-crawler-freelance-test/internal/domain/entity/user"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (authHandler *AuthHandler) Register() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
		if err := e.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
			})
		}

		if err := e.Validate(user); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					} else if e.Tag() == "email" {
						message = "Invalid email address"
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": message,
				})
			}
		}

		err := authHandler.authUsecase.Register(user)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Register Successfull",
		})
	}
}

func (authHandler *AuthHandler) Login() echo.HandlerFunc {
	return func(e echo.Context) error {
		var loginUser struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}

		if err := e.Bind(&loginUser); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
			})
		}

		if err := e.Validate(loginUser); err != nil {
			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				message := ""
				for _, e := range validationErrs {
					if e.Tag() == "required" {
						message = fmt.Sprintf("%s is required", e.Field())
					} else if e.Tag() == "email" {
						message = "Invalid email address"
					}
				}
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": message,
				})
			}
		}

		user, err := authHandler.authUsecase.Login(loginUser.Email, loginUser.Password)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid email or password",
			})
		}

		token, err := md.CreateToken(int(user.ID), user.Email)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login Successfull",
			"token":   token,
		})
	}
}
