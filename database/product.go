package database

import (
	"alta/project/config"
	"alta/project/model"
)

func InputItem(products model.Product) (interface{}, error) {
	if err := config.DB.Save(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

//GET Product in shopping_cart
func GetProductInCart() (interface{}, error) {
	var shopping_carts []model.Shopping_cart
	if err := config.DB.Find(&shopping_carts).Error; err != nil {
		return nil, err
	}
	return shopping_carts, nil
}

func GetOneProduct(product_id int) (model.Product, error) {
	var product model.Product
	if err := config.DB.Find(&product, "id=?", product_id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetProductCategory(category string) (interface{}, error) {
	var products []model.Product
	if err := config.DB.Where("category=?", category).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func GetProductType(product_category, product_type string) (interface{}, error) {
	var products []model.Product
	if err := config.DB.Where("category=? AND type=?", product_category, product_type).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
