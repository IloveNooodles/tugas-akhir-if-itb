package history

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/labstack/echo/v4"
)

// TODO deploymentGroups -> deviceId to groupId -> bisa remove aja deviceId nya kalo gitu???
// TODO tambahin description di grup
func RegisterRoute(h Handler, e *echo.Echo) {
	historyRoute := e.Group("/api/v1/histories")
	adminHistoryRoute := e.Group("/admin-api/v1/histories")

	historyRoute.Use(middleware.ValidateJWT)
	historyRoute.GET("", h.V1GetAllByCompanyID)
	historyRoute.GET("/:id", h.V1GetByID)
	historyRoute.POST("", h.V1Create)

	adminHistoryRoute.Use(middleware.ValidateAdminAPIKey)
	adminHistoryRoute.GET("", h.V1AdminGetAll)
}
