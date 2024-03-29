package model

type Coupon struct {
	Code           string `gorm:"primaryKey;not null;type:varchar(5)" validate:"required,alphanum,len=5" json:"code"`
	Discount       uint   `gorm:"check:discount >= 1 AND discount <= 100;not null" validate:"required" json:"discount"`
	MinBasketValue uint   `gorm:"check:min_basket_value >= 0;not null" validate:"required" json:"min_basket_value"`
}
