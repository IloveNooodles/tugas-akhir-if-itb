package groups

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	groupsRoute := e.Group("/api/v1/groups")
	adminGroupsRoute := e.Group("/admin-api/v1/groups")

	groupsRoute.Use(middleware.ValidateJWT)
	groupsRoute.GET("", h.V1GetAllByCompanyID)
	groupsRoute.POST("", h.V1Create)
	groupsRoute.GET("/:id", h.V1GetByID)
	groupsRoute.DELETE("/:id", h.V1Delete)
	groupsRoute.GET("/:id/devices", h.V1GetByDeviceID)

	adminGroupsRoute.Use(middleware.ValidateAdminAPIKey)
	adminGroupsRoute.GET("", h.V1AdminGetAll)
}
