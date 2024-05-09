package user

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(h Handler, e *echo.Echo) {
	usersRoute := e.Group("/api/v1/users")
	adminUsersRoute := e.Group("/admin-api/v1/users")

	usersRoute.POST("", h.V1Create)
	usersRoute.POST("/login", h.V1Login)
	usersRoute.POST("/refresh", h.V1Refresh)
	usersRoute.GET("", h.V1GetAll, middleware.ValidateJWT)
	usersRoute.GET("/:id", h.V1GetByID, middleware.ValidateJWT)
	usersRoute.DELETE("/:id", h.V1Delete, middleware.ValidateJWT)

	adminUsersRoute.GET("", h.V1AdminGetAll, middleware.ValidateAdminAPIKey, middleware.ValidateJWT)
}
