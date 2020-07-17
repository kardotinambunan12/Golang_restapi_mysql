package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)
func MasterUser (g *echo.Group){
	DEFINE_URL :="/users"
	g.GET(DEFINE_URL+"/get_data/",controllers.GetUsersController)
	g.POST(DEFINE_URL+"/add_data/",controllers.AddUserController)
	g.POST(DEFINE_URL+"/edit_data/:id/",controllers.EditUsersController)
	g.POST(DEFINE_URL+"/delete_data/:id/", controllers.DeleteUsersController)

}
