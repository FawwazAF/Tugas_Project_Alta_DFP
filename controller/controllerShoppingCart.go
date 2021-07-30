package controller

import (
	"alta/project/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateShoppingCartController(c echo.Context) error {
	// checking id_user
	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id user",
		})
	}
	// checking id_product
	id_product, err := strconv.Atoi(c.Param("id_product"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id product",
		})
	}
	// select product by id_product
	getProduct, err := database.GetOneProduct(id_product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data product by id",
		})
	}
	// insert product into cart
	newShoppingCart, err := database.InsertProductIntoCart(id_user, id_product, getProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":        "success create new shopping cart",
		"shopping cart": newShoppingCart,
	})
}
