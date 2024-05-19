package handler

import "github.com/labstack/echo/v4"

func RegisterRoute(e *echo.Echo, h Handler) {
	commonGroup := e.Group("/api/v1")

	commonGroup.GET("/health", h.PingHandler)
}
