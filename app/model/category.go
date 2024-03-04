package model

import "strings"

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

func (s *Category) Seed() *[]Category {
	contents := []string{
		"food|Food",
		"cloth|Cloth",
		"tool|Tool",
	}

	c := []Category{}
	for _, content := range contents {
		data := strings.Split(content, "|")
		c = append(c, Category{
			CategoryAPI: CategoryAPI{
				Code: &data[0],
				Name: &data[1],
			},
		})
	}
	return &c
}
