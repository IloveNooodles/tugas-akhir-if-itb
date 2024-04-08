package user

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	usersRoute := e.Group("/api/v1/users")
	usersRoute.Use(middleware.ValidateAPIKey)

	usersRoute.POST("", h.V1Create)
	usersRoute.POST("/login", h.V1Login)
	usersRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
}
