package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
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

func handleDeleteRequest(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE request")
}

func main() {
	fmt.Println("arter tendean")
	e := echo.New()

	e.GET("/", handleGetRequest)
	e.POST("/", handlePostRequest)
	e.PUT("/", handlePutRequest)
	e.DELETE("/", handleDeleteRequest)

	e.Logger.Fatal(e.Start(":8000"))
}
