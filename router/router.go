package routes

import (
	"alta/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	//Basic Routes for Skeleton **Dont Touch**
	e.GET("/users", controller.GetManyController)
	e.POST("/users", controller.DummyController)
	e.POST("/products", controller.NewItem)

	//Work your code here

	//GET show all product by category and type (Mba patmiza)

	//POST Add product to shopping cart (Mas Doni)

	//GET product form shopping cart (Fawwaz)

}
