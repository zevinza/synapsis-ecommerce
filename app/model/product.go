package model

import "github.com/google/uuid"

// Product Product
type Product struct {
	Base
	DataOwner
	ProductAPI
}

// ProductAPI Product API
type ProductAPI struct {
	SKU         *string    `json:"sku,omitempty" example:"PR-0000001" gorm:"type:varchar(256);index:idx_products_SKU_unique,unique,where:deleted_at is null;not null"` // SKU
	Name        *string    `json:"name,omitempty" example:"Sepatu" gorm:"type:varchar(256);not null"`                                                                  // Name
	Description *string    `json:"description,omitempty" gorm:"type:text"`                                                                                             // Description
	Price       *float64   `json:"price,omitempty" example:"10000" gorm:"not null"`                                                                                    // Price
	CategoryID  *uuid.UUID `json:"categoryid,omitempty" swaggertype:"string" format:"uuid"`                                                                            // CategoryID
}
