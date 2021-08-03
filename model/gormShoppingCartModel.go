package model

import "gorm.io/gorm"

type GormShoppingCartModel struct {
	DB *gorm.DB
}

func (m *GormShoppingCartModel) Get() []Shopping_cart {
	var carts []Shopping_cart
	m.DB.Find(&carts)
	return carts
}

func (m *GormShoppingCartModel) Insert(cart Shopping_cart) error {
	tx := m.DB.Save(&cart)
	return tx.Error
}

func NewGormShoppingCartModel(DB *gorm.DB) *GormShoppingCartModel {
	return &GormShoppingCartModel{DB: DB}
}
