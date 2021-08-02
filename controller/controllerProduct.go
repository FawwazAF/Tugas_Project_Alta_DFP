package controller

import (
	"alta/project/database"
	_ "alta/project/middleware"
	"net/http"

	"github.com/labstack/echo"
)

//Work Your Code Here

func ProductInCart(c echo.Context) error {
	// user_id, err := strconv.Atoi(c.Param("userid"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "please login first",
	// 	})
	// }
	product, err := database.GetProductInCart()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product in cart",
		"user":    product,
	})
}
