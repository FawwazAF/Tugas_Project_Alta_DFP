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
