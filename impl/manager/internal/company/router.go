package company

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	companiesRoute := e.Group("/api/v1/companies")
	adminCompaniesRoute := e.Group("/admin-api/v1/companies")

	companiesRoute.POST("", h.V1Create)
	companiesRoute.GET("", h.V1GetByID, middleware.ValidateJWT)

	adminCompaniesRoute.GET("", h.V1AdminGetAll, middleware.ValidateAdminAPIKey, middleware.ValidateJWT)
}
