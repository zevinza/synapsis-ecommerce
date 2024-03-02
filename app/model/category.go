package model

// Category Category
type Category struct {
	Base
	DataOwner
	CategoryAPI
}

// CategoryAPI Category API
type CategoryAPI struct {
	Code *string `json:"code,omitempty" example:"food" gorm:"type:varchar(256);index:idx_categories_code_unique,unique,where:deleted_at is null;not null"` // Code
	Name *string `json:"name,omitempty" example:"Food" gorm:"type:varchar(256)"`                                                                           // Name
}
