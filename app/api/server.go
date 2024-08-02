//go:build !js

package api

import (
	"github.com/Sourjaya/localHTMX/app/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServerDelay)

	e.GET("/", func(c echo.Context) error {
		componnets := components.Counter()
		return componnets.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Static("/", "../../public/")
	e.Logger.Fatal(e.Start(":4000"))
}
