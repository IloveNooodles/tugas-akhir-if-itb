package repositories

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	repositoryGroup := e.Group("/api/v1/repositories")

	repositoryGroup.Use(middleware.ValidateJWT)
	repositoryGroup.GET("", h.V1AdminGetAll)
	repositoryGroup.POST("", h.V1Create)
	repositoryGroup.GET("/:id", h.V1GetByID)
	repositoryGroup.DELETE("/:id", h.V1Delete)
}
