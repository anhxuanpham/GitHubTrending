package router

import (
	"GitHubTrending/handler"
	middleware "GitHubTrending/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())

}
