package binding

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SearchRequest struct {
	Keyword string `query:"keyword"`
	Page    int    `query:"page"`
}

type CPO struct {
	Name string    `json:"name"`
	Item []CPOItem `json:"item"`
}

type CPOItem struct {
	Name string `json:"name"`
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

func HandleSearchRequest(c echo.Context) error {
	searchRequest := new(SearchRequest)

	if err := c.Bind(searchRequest); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if searchRequest.Page <= 0 {
		searchRequest.Page = 1
	}

	fmt.Printf("searchRequest: %+v\n", searchRequest)

	if searchRequest.Keyword == "" {
		return c.String(http.StatusBadRequest, "keyword is required")
	}

	return c.String(http.StatusOK, "Searchcing for "+searchRequest.Keyword+" on page "+strconv.Itoa(searchRequest.Page))
}

func HandleCPO(c echo.Context) error {
	cpo := new(CPO)

	if err := c.Bind(cpo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if cpo.Name == "" {
		return c.String(http.StatusBadRequest, "CPO name is required")
	}

	if cpo.Item == nil {
		cpo.Item = []CPOItem{}
	}

	for _, cpoi := range cpo.Item {
		if cpoi.Name == "" {
			return c.String(http.StatusBadRequest, "CPO Item name is required")
		}
	}

	return c.JSON(http.StatusOK, cpo)
}
