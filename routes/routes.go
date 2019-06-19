package routes

import (
	c "teak/controllers"
	"teak/config"
	teakMiddleware "teak/modules/teak/middleware"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func acl(permission string) echo.MiddlewareFunc {
	return teakMiddleware.ACL(permission)
}

func Api(e *echo.Echo) {
	e.POST("login", c.LoginUser)
	e.GET("profile", c.GetProfile, middleware.JWTWithConfig(config.JwtConfig))

	// user
	user := e.Group("user", middleware.JWTWithConfig(config.JwtConfig))
	user.GET("", c.GetUser)
	user.POST("", c.AddUser, acl("your_permission"))
	user.PUT("/:id", c.UpdateUser, acl("your_permission"))
	
}
