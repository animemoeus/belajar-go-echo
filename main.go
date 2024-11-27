package main

import (
	"belajar-echo/binding"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func handleGetRequest(c echo.Context) error {
	return c.String(http.StatusOK, "arter tendean")
}

func handlePostRequest(c echo.Context) error {
	return c.String(http.StatusOK, "arter")
}

func handlePutRequest(c echo.Context) error {
	return c.String(http.StatusOK, "PUT request")
}

type User struct {
	Name    string
	Address string
	Age     int
}

func handleDeleteRequest(c echo.Context) error {
	//return c.String(http.StatusOK, "DELETE request")
	user := User{"arter tendean", "jakarta", 123}
	return c.JSON(http.StatusOK, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	fmt.Println("id:", id)

	if id == "" {
		return c.String(http.StatusOK, "`id` params is missing")
	}

	return c.String(http.StatusOK, "id params: "+id)
}

func main() {
	fmt.Println("arter tendean")
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", handleGetRequest)
	e.POST("/", handlePostRequest)
	e.PUT("/", handlePutRequest)
	e.DELETE("/", handleDeleteRequest)

	e.GET("/users/:id", getUser)

	e.POST("/binding/register", binding.HandleUserRegistration)
	e.GET("/binding/search", binding.HandleSearchRequest)
	e.POST("/binding/cpo", binding.HandleCPO)

	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
