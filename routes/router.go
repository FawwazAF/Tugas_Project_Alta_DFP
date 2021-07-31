package routes

import (
	"alta/project/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	//Basic Routes for Skeleton **Dont Touch**
	e.GET("/users", controller.GetManyController)
	e.POST("/products", controller.NewItem)
	e.DELETE("/users/:id", controller.DeleteUser)

	//Work your code here

	//GET show all product by category and type (Mba patmiza)

	//POST Add product to shopping cart (Mas Doni)
	e.POST("carts/:user_id/:product_id", controller.CreateShoppingCartController)
	//DELETE product to shopping cart (Mas Doni)
	e.DELETE("carts/:user_id/:product_id", controller.DeleteShoppingCartController)

	//Fawwaz Workspace
	//GET product form shopping cart
	e.GET("/cart", controller.ProductInCart)

	//Register User
	e.POST("/users", controller.UserRegister)

	//Login
	e.POST("/login", controller.LoginUser)

}
