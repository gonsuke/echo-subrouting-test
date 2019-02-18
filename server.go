package main

import (
	"net/http"

	"github.com/labstack/echo"
)

const Hello = "Hello, World!"

func main() {
	router := NewRouter()
	router.Start(":8080")
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/hello", handler)

	subRoute1 := e.Group("/")
	subRoute1.GET("/hello2", handler) // This handler will NOT be found.

	subRoute2 := e.Group("/subroute2")
	subRoute2.GET("/hello3", handler) // This handler will be found.
	return e
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, Hello)
}
