package groups

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	groupsRoute := e.Group("/api/v1/groups")
	adminGroupsRoute := e.Group("/admin-api/v1/groups")

	groupsRoute.POST("", h.V1Create)
	groupsRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
	groupsRoute.GET("/:id/devices", h.V1GetByDeviceID, middleware.ValidateJWT)

	adminGroupsRoute.GET("", h.V1AdminGetAll)
}
