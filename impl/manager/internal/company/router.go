package company

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	companiesRoute := e.Group("/api/v1/companies")
	companiesRoute.Use(middleware.ValidateJWT)
	companiesRoute.GET("", h.V1GetCompanyAndLoggedInUser)
	companiesRoute.GET("/server/health", h.V1CheckClusterStatus)

	adminCompaniesRoute := e.Group("/admin-api/v1/companies")
	adminCompaniesRoute.Use(middleware.ValidateAdminAPIKey)
	adminCompaniesRoute.POST("", h.V1Create)
	adminCompaniesRoute.GET("", h.V1AdminGetAll)
	adminCompaniesRoute.GET("/:id", h.V1GetByID)
	adminCompaniesRoute.DELETE("/:id", h.V1Delete)
}
