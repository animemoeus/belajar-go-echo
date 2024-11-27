package binding

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandleUserRegistration(c echo.Context) error {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if registerRequest.Username == "" {
		return c.String(http.StatusBadRequest, "username is required")
	}

	if registerRequest.Password == "" {
		return c.String(http.StatusBadRequest, "password is required")
	}

	if registerRequest.Email == "" {
		return c.String(http.StatusBadRequest, "email is required")
	}

	return c.String(http.StatusOK, "User "+registerRequest.Username+" has been registered")
}
