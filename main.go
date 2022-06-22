package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", serveBudge)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func serveBudge(c echo.Context) error {
	f, err := os.Open("assets/budge.svg")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/svg+xml", f)

	// return c.String(http.StatusOK, "Hello, World!")
}
