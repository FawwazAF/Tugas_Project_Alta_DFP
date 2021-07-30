package database

import (
	"alta/project/config"
	"alta/project/model"
)

func InsertProductIntoCart(id_user, id_product int, product model.Product) (interface{}, error) {
	// insert and select data product to cart
	shoppingCart := model.Shopping_cart{
		User_id:    id_user,
		Product_id: id_product,
		Name:       product.Name,
		Category:   product.Category,
		Type:       product.Type,
		Price:      product.Price,
		Qty:        1,
	}
	if err := config.DB.Save(&shoppingCart).Error; err != nil {
		return shoppingCart, err
	}
	return shoppingCart, nil
}
