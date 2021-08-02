package model

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	Total_qty   int
	Total_price int
}
