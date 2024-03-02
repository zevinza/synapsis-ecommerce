package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Business model
type Business struct {
	Base
	DataOwner
	BusinessAPI
}

// BusinessAPI model
type BusinessAPI struct {
	OwnerID        *uuid.UUID       `json:"owner_id,omitempty" gorm:"index;type:uuid"`
	BusinessName   *string          `json:"business_name,omitempty" validate:"omitempty,emptyString"`
	JoinDate       *strfmt.DateTime `json:"join_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	IsActive       *bool            `json:"is_active,omitempty"`
	ActivatedAt    *strfmt.DateTime `json:"activated_at,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	IsSyncronize   *bool            `json:"is_syncronize,omitempty"`
	ActivePackage  *uuid.UUID       `json:"active_package,omitempty"`
	PackageEndDate *strfmt.DateTime `json:"package_end_date,omitempty" swaggerignore:"true"`
}
