package model

type Basket struct {
	Value                 uint `json:"value"`
	AppliedDiscount       uint `json:"applied_discount"`
	ApplicationSuccessful bool `json:"application_successful"`
}
