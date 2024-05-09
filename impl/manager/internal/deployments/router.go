package deployments

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	deploymentsRoute := e.Group("/api/v1/deployments")
	admindeploymentsRoute := e.Group("/admin-api/v1/deployments")

	deploymentsRoute.Use(middleware.ValidateJWT)
	deploymentsRoute.GET("", h.V1GetAllByCompanyID)
	deploymentsRoute.GET("/:id", h.V1GetByID)
  deploymentsRoute.DELETE("/:id", h.V1Delete)
	deploymentsRoute.POST("", h.V1Create)

	deployRoutes := deploymentsRoute.Group("/deploy")
	deployRoutes.POST("", h.V1Deploy)
	deployRoutes.POST("/delete", h.V1DeleteDeploy)

	admindeploymentsRoute.Use(middleware.ValidateAdminAPIKey)
	admindeploymentsRoute.GET("", h.V1AdminGetAll)
}
