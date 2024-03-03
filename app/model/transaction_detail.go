package model

import "github.com/google/uuid"

type TransactionDetail struct {
	Base
	DataOwner
	TransactionDetailAPI
}
type TransactionDetailAPI struct {
	TransactionID      *uuid.UUID `json:"transaction_id,omitempty" swaggertype:"string" format:"uuid"`
	ProductID          *uuid.UUID `json:"product_id,omitempty" swaggertype:"string" format:"uuid"`
	ProductName        *string    `json:"product_name,omitempty"`
	ProductSKU         *string    `json:"product_sku,omitempty"`
	ProductDescription *string    `json:"product_description,omitempty"`
	ProductPrice       *float64   `json:"product_price,omitempty"`
	Quantity           *int64     `json:"quantity,omitempty"`
	SubTotalPrice      *float64   `json:"sub_total_price,omitempty"`
	Notes              *string    `json:"notes,omitempty"`
}
