package device

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	devicesRoute := e.Group("/api/v1/devices")
	adminDevicesRoute := e.Group("/admin-api/v1/devices")

	devicesRoute.POST("", h.V1Create)
	devicesRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
	devicesRoute.GET("/:id/groups", h.V1GetByGroupID, middleware.ValidateJWT)

	adminDevicesRoute.GET("", h.V1AdminGetAll, middleware.ValidateAdminAPIKey, middleware.ValidateJWT)
}
