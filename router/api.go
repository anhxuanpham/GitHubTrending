package router

import (
	"GitHubTrending/handler"
	myMiddleware "GitHubTrending/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, myMiddleware.IsAdmin())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

}
