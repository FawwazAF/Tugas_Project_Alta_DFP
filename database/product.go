package database

import (
	"alta/project/config"
	"alta/project/model"
)

func GetProducts() (interface{}, error) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func InputItem(products model.Product) (interface{}, error) {
	if err := config.DB.Save(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
