package model

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model            // ID, CreatedAt, DeletedAt, UpdatedAt
	Code           string `gorm:"not null" json:"code"`
	Discount       uint   `gorm:"check:Discount >= 1 AND Discount <= 100;not null" json:"discount"`
	MinBasketValue uint   `gorm:"check:MinBasketValue >= 1;not null" json:"min_basket_value"`
}

/*
We use `uint` but PostgreSQL
doesn't have unsigned type for integers.
*/
