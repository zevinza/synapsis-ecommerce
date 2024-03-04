package model

import (
	"time"

	"github.com/google/uuid"
)

// LoginAPI model
type LoginAPI struct {
	Username *string `json:"username,omitempty" example:"armadamuhammads@gmail.com" validate:"required"`
	Password *string `json:"password,omitempty" example:"@Password123" validate:"required"`
	Remember *bool   `json:"remember,omitempty" example:"true"`
}

// RegistrationAPI model
type RegistrationAPI struct {
	FirstName       *string `json:"first_name,omitempty" example:"Armada" validate:"required"`
	LastName        *string `json:"last_name,omitempty" example:"Muhammad"`
	Email           *string `json:"email,omitempty" example:"armadamuhammads@gmail.com" validate:"required,email"`
	Username        *string `json:"username,omitempty" example:"armadamuhammads"`
	PhoneNumber     *string `json:"phone_number,omitempty" example:"089678009400"`
	Password        *string `json:"password,omitempty" example:"@Password123" validate:"required,min=8"`
	ConfirmPassword *string `json:"confirm_password,omitempty" example:"@Password123" validate:"required,eqfield=Password"`
}

// VerificationAccountAPI model
type VerificationAccountAPI struct {
	Email            *string `json:"email,omitempty" example:"armadamuhammads@gmail.com" validate:"required"`
	VerificationCode *string `json:"verification_code,omitempty" example:"7085" validate:"required"`
}

// LoginResponse model
type LoginResponse struct {
	Token *ResponseToken `json:"token,omitempty"`
	User  *User          `json:"user,omitempty"`
}

// RegistrationResponse model
type RegistrationResponse struct {
	User *User `json:"user,omitempty"`
}

// ResponseToken model
type ResponseToken struct {
	AccessToken  *string `json:"access_token,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
	IsAdmin      *bool   `json:"is_admin,omitempty"`
}

type Auth struct {
	UserID  *uuid.UUID `json:"user_id,omitempty"`
	IsAdmin *bool      `json:"is_admin,omitempty"`
	Exp     *int64     `json:"exp,omitempty"`
}

// ResponseAuthenticate model
type ResponseAuthenticate struct {
	ClientID            *string    `json:"ClientID,omitempty"`
	UserID              *string    `json:"UserID,omitempty"`
	IsAdmin             *bool      `json:"IsAdmin,omitempty"`
	RedirectURI         *string    `json:"RedirectURI,omitempty"`
	Scope               *string    `json:"Scope,omitempty"`
	Code                *string    `json:"Code,omitempty"`
	CodeChallenge       *string    `json:"CodeChallenge,omitempty"`
	CodeChallengeMethod *string    `json:"CodeChallengeMethod,omitempty"`
	CodeCreateAt        *time.Time `json:"CodeCreateAt,omitempty"`
	CodeExpiresIn       *int       `json:"CodeExpiresIn,omitempty"`
	Access              *string    `json:"Access,omitempty"`
	AccessCreateAt      *time.Time `json:"AccessCreateAt,omitempty"`
	AccessExpiresIn     *int64     `json:"AccessExpiresIn,omitempty"`
	Refresh             *string    `json:"Refresh,omitempty"`
	RefreshCreateAt     *time.Time `json:"RefreshCreateAt,omitempty"`
	RefreshExpiresIn    *int64     `json:"RefreshExpiresIn,omitempty"`
}
