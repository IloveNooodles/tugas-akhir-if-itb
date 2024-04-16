package device

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	devicesRoute := e.Group("/api/v1/devices")

	devicesRoute.POST("", h.V1Create)
	devicesRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
}
