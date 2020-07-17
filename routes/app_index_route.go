package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func AppIndex(g *echo.Group) {
	
	g.GET("", controllers.AppIndex)
}
