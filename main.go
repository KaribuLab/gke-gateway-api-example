package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(1)
	e.GET("/", func(c echo.Context) error {
		e.Logger.Error("Unauthorized")
		return c.String(http.StatusUnauthorized, "Unauthorized")
	})
	e.GET("/health", func(c echo.Context) error {
		e.Logger.Info("Health check")
		return c.String(http.StatusOK, "Up")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
