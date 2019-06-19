package middleware

import (
	//"net/http"

	"github.com/labstack/echo"
)

// ACL is method for checking user permisson
func ACL(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Put your custom permission
			return next(c)
			// return c.JSON(http.StatusForbidden, "You don't have permission to access this resource")
		}
	}
}