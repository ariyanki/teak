package routes

import (
	c "teak/controllers"
	
	"github.com/labstack/echo"
)

func Api(e *echo.Echo) {
	// user
	e.POST("login", c.LoginUser)
}
