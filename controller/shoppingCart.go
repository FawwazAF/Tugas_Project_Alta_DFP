package controller

import (
	"alta/project/model"
	"fmt"

	"github.com/labstack/echo"
)

func CreatePostShoppingCartController(cartModel model.ShoppingCartModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cart model.Shopping_cart
		fmt.Printf("%#v", cart)
		return c.JSON(200, cart)
	}
}
