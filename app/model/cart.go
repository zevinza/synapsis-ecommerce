package model

import "github.com/google/uuid"

// Cart Cart
type Cart struct {
	Base
	DataOwner
	CartAPI
}

// CartAPI Cart API
type CartAPI struct {
	UserID    *uuid.UUID `json:"userid,omitempty" swaggertype:"string" format:"uuid"`    // UserID
	ProductID *uuid.UUID `json:"productid,omitempty" swaggertype:"string" format:"uuid"` // ProductID
	Quantity  *int64     `json:"quantity,omitempty" example:"1"`                         // Quantity
	Price     *float64   `json:"price,omitempty" example:"10000"`                        // Price
}
