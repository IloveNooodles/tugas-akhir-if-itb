package middleware

import (
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/auth"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var cfg = config.Config{}

func init() {
	cfg = config.New()
}

// Middleware to validate api Key that comes from request headers.
//
// It looks for the "X-Api-Key" from the headers and match it with current api key from .env
func ValidateAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-Api-Key")
		if apiKey != cfg.ApiKey {
			return c.JSON(http.StatusUnauthorized, handler.ErrorResponse{Message: "invalid API Key"})
		}

		return next(c)
	}
}

// Middleware to validate AdminApiKey that comes from request headers
//
// It looks for the "X-Admin-Api-Key" from the headers and match it with current api key from .env
func ValidateAdminAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-Admin-Api-Key")
		if apiKey != cfg.AdminAPIKey {
			return c.JSON(http.StatusForbidden, handler.ErrorResponse{Message: "Forbidden"})
		}

		return next(c)
	}
}

// Middleware to validate JWT Token that comes from the cookie
func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Cookie method
		cookie, err := c.Cookie("accessToken")
		if err != nil {
			log.Errorf("error when getting cookie, err: %s", err)
			return c.JSON(http.StatusUnauthorized, handler.ErrorResponse{Message: "access token not found"})
		}

		// authToken := c.Request().Header.Get("Authorization")
		validToken, err := auth.ValidateToken(cookie.Value)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, handler.ErrorResponse{Message: "invalid token"})
		}

		claims, ok := validToken.Claims.(*auth.JwtClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "server error"})
		}

		c.Set("userID", claims.UserID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("companyID", claims.CompanyID)
		return next(c)
	}
}
