package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	e := echo.New()

	//e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello From Echo",
		})
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "saludos desde Echo")

	})

	e.Start(":8080")
}
