package middleware

import (
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/dto"
	"github.com/labstack/echo/v4"
)

var cfg = config.Config{}

func init() {
	cfg = config.New()
}

func ValidateAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-Api-Key")
		if apiKey != cfg.ApiKey {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "invalid API Key"})
		}

		return next(c)
	}
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
