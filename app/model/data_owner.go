package model

import "github.com/google/uuid"

// DataOwner is an abstract struct
type DataOwner struct {
	CreatorID  *uuid.UUID `json:"creator_id,omitempty" gorm:"type:varchar(36)" swaggerignore:"true"`                  // creator id
	ModifierID *uuid.UUID `json:"modifier_id,omitempty" gorm:"type:varchar(36)" swaggerignore:"true"`                 // modifier id
	Creator    *User      `json:"creator,omitempty" gorm:"foreignKey:CreatorID;references:ID" swaggerignore:"true"`   // creator
	Modifier   *User      `json:"modifier,omitempty" gorm:"foreignKey:ModifierID;references:ID" swaggerignore:"true"` // modifier
}
