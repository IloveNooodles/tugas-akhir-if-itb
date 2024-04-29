package history

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	groupsRoute := e.Group("/api/v1/histories")
	adminGroupsRoute := e.Group("/admin-api/v1/histories")

	groupsRoute.Use(middleware.ValidateJWT)
	groupsRoute.POST("", h.V1Create)
	groupsRoute.GET("/:id", h.V1GetByID)

	adminGroupsRoute.Use(middleware.ValidateAdminAPIKey)
	adminGroupsRoute.GET("", h.V1AdminGetAll)
}
