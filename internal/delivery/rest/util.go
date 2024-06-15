package rest

import (
	"time"

	"github.com/labstack/echo/v4"
)

var spawn = time.Now()

func RegisterUtilHandler(r *echo.Group) {
	r.GET("/status", healthcheck)
	r.GET("/status", healthcheck)
}

func healthcheck(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"status": "healthy",
		"uptime": time.Since(spawn).String(),
	})
}
