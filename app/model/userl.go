package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
type User struct {
	Base
	DataOwner
	UserAPI
}

// UserAPI model
type UserAPI struct {
	AssetID                   *uuid.UUID       `json:"asset_id,omitempty" gorm:"type:varchar(36)"`
	GroupID                   *uuid.UUID       `json:"group_id,omitempty" gorm:"type:varchar(36)"`
	BusinessID                *uuid.UUID       `json:"business_id,omitempty" gorm:"type:varchar(36)"`
	EmployeeID                *uuid.UUID       `json:"employee_id,omitempty" gorm:"type:varchar(36)"`
	IsOwner                   *bool            `json:"is_owner,omitempty"`
	Fullname                  *string          `json:"fullname,omitempty" gorm:"not null"`
	Username                  *string          `json:"username,omitempty" gorm:"type:varchar(191);index:idx_user_username_unique,unique,where:deleted_at is null;not null"`
	Email                     *string          `json:"email,omitempty" gorm:"type:varchar(191);index:idx_user_email_unique,unique,where:deleted_at is null;not null"`
	Mobile                    *string          `json:"mobile,omitempty"`
	Password                  *string          `json:"-" gorm:"not null"`
	Salt                      *string          `json:"-,omitempty"`
	IsPasswordSystemGenerated *bool            `json:"is_password_system_generated,omitempty"`
	PasswordLastChange        *strfmt.DateTime `json:"password_last_change,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	PasswordExpiration        *strfmt.DateTime `json:"password_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	ResetPasswordCode         *string          `json:"-"`
	ResetPasswordExpiration   *strfmt.DateTime `json:"reset_password_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	OTPEnabled                *bool            `json:"otp_enabled,omitempty"`
	OTPCode                   *string          `json:"-"`
	OTPExpiration             *strfmt.DateTime `json:"otp_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	VerificationCode          *string          `json:"-"`
	IsVerified                *bool            `json:"is_verified,omitempty"`
	VerificationExpiration    *strfmt.DateTime `json:"verification_expiration,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	LastLogin                 *strfmt.DateTime `json:"last_login,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	LoginAttempt              *int             `json:"login_attempt,omitempty"`
}

// BeforeCreate Data
func (b *User) BeforeCreate(tx *gorm.DB) error {
	if nil != b.ID {
		return nil
	}
	id, e := uuid.NewRandom()
	now := strfmt.DateTime(time.Now())
	b.ID = &id
	b.CreatedAt = &now
	b.UpdatedAt = &now

	// creating username
	var count int64
	var email string
	if b.Email != nil {
		email = *b.Email
	}
	prefixEmail := strings.Split(email, "@")
	s := prefixEmail[0] + "@%"
	user := User{}
	tx.Model(&user).Where(`email LIKE ?`, s).Count(&count)
	username := prefixEmail[0]
	if count > 0 {
		username = prefixEmail[0] + fmt.Sprint(count)
	}
	b.Username = &username

	return e
}

// UserData
type UserData struct {
	Username *string `json:"username,omitempty" example:"john.doe" gorm:"type:varchar(191);index:idx_user_username_unique,unique,where:deleted_at is null;not null"`
	Password *string `json:"password" example:"password"  gorm:"not null"`
}

type UpdateProfileAPI struct {
	AssetID          *uuid.UUID   `json:"asset_id,omitempty" gorm:"type:varchar(36)"`
	Fullname         *string      `json:"fullname,omitempty" example:"John Doe"`                                                                               // Fullname
	Email            *string      `json:"email,omitempty" example:"john.doe@mail.com" validate:"email"`                                                        // Email
	Username         *string      `json:"username,omitempty" example:"johndoe"`                                                                                // Username
	Mobile           *string      `json:"mobile,omitempty" example:"08123456789" validate:"omitempty,phone"`                                                   // Mobile
	AlternateNumber  *string      `json:"alternate_number,omitempty" example:"08123456789" gorm:"type:varchar(191)" validate:"omitempty,phone"`                // Alternate NumberAssetID          *uuid.UUID   `json:"asset_id,omitempty" swaggertype:"string" format:"uuid"`                                  // AssetID
	IsChangePassword *bool        `json:"is_change_password,omitempty" example:"true"`                                                                         // IsChangePassword
	OldPassword      *string      `json:"old_password,omitempty" example:"@Password123" validate:"required_if=IsChangePassword true"`                          // OldPassword
	NewPassword      *string      `json:"new_password,omitempty" example:"@Password1234" validate:"required_if=IsChangePassword true"`                         // NewPassword
	ConfirmPassword  *string      `json:"confirm_password,omitempty" example:"@Password1234" validate:"required_if=IsChangePassword true,eqfield=NewPassword"` // ConfirmPassword
	Gender           *string      `json:"gender,omitempty" example:"male" validate:"omitempty,oneof=male female 'rather_not_say'"`                             // Gender : male | female | rather_not_say
	MaritalStatus    *string      `json:"marital_status,omitempty" example:"married" validate:"omitempty,oneof=married unmarried divorced"`                    // MaritalStatus : married | unmarried | divorced
	DateOfBirth      *strfmt.Date `json:"date_of_birth,omitempty" example:"2000-01-02" swaggerignore:"true"`                                                   // DateOfBirth
	ProvinceID       *int64       `json:"province_id,omitempty"`                                                                                               // ProvinceID
	CityID           *int64       `json:"city_id,omitempty"`                                                                                                   // CityID
	SubdistrictID    *int64       `json:"subdistrict_id,omitempty"`                                                                                            // SubdistrictID
	Address          *string      `json:"address,omitempty" example:"Jl. Pegangsaan Timur No. 56"`                                                             // Address
	ZipCode          *string      `json:"zip_code,omitempty" example:"57588"`                                                                                  // Address
}

type UserPayload struct {
	AssetID     *uuid.UUID `json:"asset_id,omitempty" gorm:"type:varchar(36)"`
	GroupID     *uuid.UUID `json:"group_id,omitempty" gorm:"type:varchar(36)"`
	EmployeeID  *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36)"`
	Fullname    *string    `json:"fullname,omitempty" gorm:"not null"`
	Email       *string    `json:"email,omitempty" gorm:"type:varchar(191);index:idx_user_email_unique,unique,where:deleted_at is null;not null"`
	Mobile      *string    `json:"mobile,omitempty"`
	Password    *string    `json:"password" gorm:"not null"`
	IsActivated *bool      `json:"is_activated,omitempty"`
}

type UserUpdate struct {
	GroupID              *uuid.UUID `json:"group_id,omitempty" gorm:"type:varchar(36)"`
	Status               *string    `json:"status,omitempty"`
	Username             *string    `json:"username,omitempty"`
	IsPasswordUpdate     *bool      `json:"is_password_update,omitempty"`
	Password             *string    `json:"password,omitempty"`
	PasswordConfirmation *string    `json:"password_confirmation,omitempty"`
}
