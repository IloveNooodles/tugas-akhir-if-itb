package middleware

import (
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/auth"
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

func ValidateAdminAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-Admin-Api-Key")
		if apiKey != cfg.AdminAPIKey {
			return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Forbidden"})
		}

		return next(c)
	}
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Cookie method
		var authToken = ""
		cookies := c.Cookies()
		for _, ck := range cookies {
			if ck.Name == "accessToken" {
				authToken = ck.Value
			}
		}

		// authToken := c.Request().Header.Get("Authorization")
		validToken, err := auth.ValidateToken(authToken)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "invalid token"})
		}

		claims, ok := validToken.Claims.(*auth.JwtClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "server error"})
		}
		c.Set("userID", claims.UserID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("companyID", claims.CompanyID)
		return next(c)
	}
}
