package company

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	companiesRoute := e.Group("/api/v1/companies")
	companiesRoute.Use(middleware.ValidateAPIKey)

	companiesRoute.POST("", h.V1Create)
	companiesRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
}
