package device

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	devicesRoute := e.Group("/api/v1/devices")
	adminDevicesRoute := e.Group("/admin-api/v1/devices")

	devicesRoute.Use(middleware.ValidateJWT)
	devicesRoute.GET("", h.V1GetAllByCompanyID)
	devicesRoute.POST("", h.V1Create)
	devicesRoute.GET("/:id", h.V1GetByID)
  devicesRoute.DELETE("/:id", h.V1Delete)
	devicesRoute.GET("/:id/groups", h.V1GetGroupByDeviceID)

	adminDevicesRoute.Use(middleware.ValidateAdminAPIKey)
	adminDevicesRoute.GET("", h.V1AdminGetAll)
}
