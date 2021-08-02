package database

import (
	"alta/project/config"
	"alta/project/model"
)

//GET total price and Qty
func GetTotal() (model.Checkout, error) {
	var shopping_carts []model.Shopping_cart
	var checkout model.Checkout
	if err := config.DB.Find(&shopping_carts).Error; err != nil {
		return checkout, err
	}
	total_qty := 0
	total_price := 0
	for _, v := range shopping_carts {
		total_qty += v.Qty
		total_price = total_price + (v.Qty * v.Price)
	}
	checkout = model.Checkout{
		Total_qty:   total_qty,
		Total_price: total_price,
	}
	if err := config.DB.Save(&checkout).Error; err != nil {
		return checkout, err
	}
	return checkout, nil
}
