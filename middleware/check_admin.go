package middleware

import (
	"GitHubTrending/model"
	"GitHubTrending/model/req"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {

				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "admin@gmai.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "You are not permission call this API",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}
