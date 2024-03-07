package model

import "github.com/google/uuid"

// Cart Cart
type Cart struct {
	Base
	DataOwner
	CartAPI
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ID;references:ProductID"`
}

// CartAPI Cart API
type CartAPI struct {
	UserID    *uuid.UUID `json:"user_id,omitempty" gorm:"type:varchar(256);index:idx_cart_user_product_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`    // UserID
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(256);index:idx_cart_user_product_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // ProductID
	Quantity  *int64     `json:"quantity,omitempty" example:"1"`                                                                                                                               // Quantity
	Price     *float64   `json:"price,omitempty" example:"10000"`                                                                                                                              // Price
	Notes     *string    `json:"notes,omitempty"`
}

type CartPayload struct {
	ProductID *uuid.UUID `json:"product_id,omitempty" swaggertype:"string" format:"uuid"` // ProductID
	Quantity  *int64     `json:"quantity,omitempty" example:"1"`                          // Quantity
	Notes     *string    `json:"notes,omitempty"`
}
