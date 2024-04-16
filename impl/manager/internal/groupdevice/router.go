package groupdevice

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	groupDeviceRoute := e.Group("/api/v1/groupdevices")
	adminGroupDeviceRoute := e.Group("/admin-api/v1/groupdevices")

	groupDeviceRoute.POST("", h.V1Create)
	groupDeviceRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)

	adminGroupDeviceRoute.GET("", h.V1AdminGetAll, middleware.ValidateAdminAPIKey, middleware.ValidateJWT)
}
