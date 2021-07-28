package database

import (
	"alta/project/config"
	_ "alta/project/middleware"
	"alta/project/model"
)

func GetUsers() (interface{}, error) {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
