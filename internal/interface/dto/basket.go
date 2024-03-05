package dto

type BasketRequestDTO struct {
	Code  string `json:"code" validate:"required"`
	Value uint   `json:"value" validate:"required,min=0"`
}

type BasketResponseDTO struct {
	Value                 uint  `json:"value"`
	AppliedDiscount       uint  `json:"applied_discount"`
	ApplicationSuccessful bool  `json:"application_successful"`
	FinalValue            *uint `json:"final_value,omitempty"`
}
