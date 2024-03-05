package model

import "github.com/go-openapi/strfmt"

// User User
type User struct {
	Base
	DataOwner
	UserAPI
}

// UserAPI User API
type UserAPI struct {
	FirstName   *string          `json:"first_name,omitempty" example:"Armada" gorm:"type:varchar(256);not null"`                                                                            // FirstName
	LastName    *string          `json:"last_name,omitempty" example:"Muhammad Siswanto" gorm:"type:varchar(256)"`                                                                           // LastName
	Username    *string          `json:"username,omitempty" gorm:"type:varchar(256);index:idx_users_Username_unique,unique,where:deleted_at is null;not null"`                               // Username
	Email       *string          `json:"email,omitempty" example:"armadamuhammads@gmail.com" gorm:"type:varchar(256);index:idx_users_email_unique,unique,where:deleted_at is null;not null"` // Email
	PhoneNumber *string          `json:"phone_number,omitempty" example:"089678009400" gorm:"type:varchar(15)"`                                                                              // PhoneNumber
	IsAdmin     *bool            `json:"is_admin,omitempty"`                                                                                                                                 // IsAdmin
	Password    *string          `json:"-" gorm:"type:text"`                                                                                                                                 // Password
	LastLogin   *strfmt.DateTime `json:"last_login,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                                                               // LastLogin
}

type UserPayload struct {
	FirstName   *string `json:"first_name,omitempty" example:"Armada" validate:"required"`
	LastName    *string `json:"last_name,omitempty" example:"Muhammad"`
	Email       *string `json:"email,omitempty" example:"armadamuhammads@gmail.com" validate:"required,email"`
	Username    *string `json:"username,omitempty" example:"armadamuhammads"`
	PhoneNumber *string `json:"phone_number,omitempty" example:"089678009400"`
}

func (s *User) Seed() *[]User {

	u := []User{
		{
			UserAPI: UserAPI{
				FirstName: strptr("Admin"),
				Username:  strptr("admin"),
				Email:     strptr("admin@mail.com"),
				IsAdmin:   boolInt(1),
				Password:  strptr("$2a$10$8katy8Li/HTJ.LljINi3oOrEIrvL.iscnzjnqskWONkJlyPxLq9W."),
				LastLogin: now(),
			},
		},
		{
			UserAPI: UserAPI{
				FirstName:   strptr("Armada"),
				LastName:    strptr("Muhammad"),
				Username:    strptr("armada_muhammad"),
				Email:       strptr("armadamuhammads@gmail.com"),
				PhoneNumber: strptr("089678009400"),
				IsAdmin:     boolInt(0),
				Password:    strptr("$2a$10$8katy8Li/HTJ.LljINi3oOrEIrvL.iscnzjnqskWONkJlyPxLq9W."),
				LastLogin:   now(),
			},
		},
	}

	return &u
}
