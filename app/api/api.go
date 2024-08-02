package api

import (
	"time"

	"github.com/labstack/echo/v4"
)

func ServerDelay(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		time.Sleep(1000 * time.Millisecond)
		return next(c)
	}
}
