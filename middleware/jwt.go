package middleware

import (
	"GitHubTrending/model"
	"GitHubTrending/security"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: security.SECRET_KEY,
	}
	return middleware.JWTWithConfig(config)
}
