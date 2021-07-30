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

func GetOneProduct(id_product int) (model.Product, error) {
	var product model.Product
	if err := config.DB.Find(&product, "id_product=?", id_product).Error; err != nil {
		return product, err
	}
	return product, nil
}
