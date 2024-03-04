package model

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product Product
type Product struct {
	Base
	DataOwner
	ProductAPI
}

// ProductAPI Product API
type ProductAPI struct {
	CategoryID   *uuid.UUID `json:"category_id,omitempty" swaggertype:"string" format:"uuid"`                                                                           // CategoryID
	SKU          *string    `json:"sku,omitempty" example:"PR-0000001" gorm:"type:varchar(256);index:idx_products_sku_unique,unique,where:deleted_at is null;not null"` // SKU
	Name         *string    `json:"name,omitempty" example:"Sepatu" gorm:"type:varchar(256);not null"`                                                                  // Name
	Description  *string    `json:"description,omitempty" gorm:"type:text"`                                                                                             // Description
	Price        *float64   `json:"price,omitempty" example:"10000" gorm:"not null"`                                                                                    // Price
	CategoryCode *string    `json:"-"`
}

func (s *Product) BeforeCreate(tx *gorm.DB) error {
	if s.CategoryCode != nil && s.CategoryID == nil {
		category := Category{}
		tx.Where(`code = ?`, s.CategoryCode).Take(&category)
		s.CategoryID = category.ID
	}

	return s.Base.BeforeCreate(tx)
}

func (s *Product) Seed() *[]Product {
	contents := []string{
		"PR-00001|Susu Bubuk|susu bubuk 100gr|5000|food",
		"PR-00002|Roti Tawar|Roti tawar potong 250gr|10000|food",
		"PR-00003|Masako Ayam|Bumbu Penyedap 100gr|2000|food",
		"PR-00004|Kaos Polos|hitam|35000|cloth",
		"PR-00005|Sepatu Karet|hitam|125000|cloth",
		"PR-00006|Topi| hitam|20000|cloth",
	}

	c := []Product{}
	for _, content := range contents {
		data := strings.Split(content, "|")

		f, _ := strconv.ParseFloat(data[3], 64)
		c = append(c, Product{
			ProductAPI: ProductAPI{
				SKU:          &data[0],
				Name:         &data[1],
				Description:  &data[2],
				Price:        &f,
				CategoryCode: &data[4],
			},
		})
	}
	return &c
}
