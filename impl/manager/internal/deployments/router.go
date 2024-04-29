package deployments

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	deploymentsRoute := e.Group("/api/v1/deployments")
	admindeploymentsRoute := e.Group("/admin-api/v1/deployments")

	deploymentsRoute.Use(middleware.ValidateJWT)
	deploymentsRoute.POST("", h.V1Create)
	deploymentsRoute.GET("/:id", h.V1GetByID)

	admindeploymentsRoute.Use(middleware.ValidateAdminAPIKey)
	admindeploymentsRoute.GET("", h.V1AdminGetAll)
}
