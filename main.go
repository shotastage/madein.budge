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
	e.Logger.Fatal(e.Start(":1323"))
}

func serveBudge(c echo.Context) error {
	f, err := os.Open("assets/budge.png")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/png", f)

	// return c.String(http.StatusOK, "Hello, World!")
}
