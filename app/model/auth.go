package model

import "time"

// LoginAPI model
type LoginAPI struct {
	Username *string `json:"username,omitempty" example:"john.doe@mail.com" validate:"required"`
	Password *string `json:"password,omitempty" example:"@Password123" validate:"required"`
	Remember *bool   `json:"remember,omitempty" example:"true"`
}

// RegistrationAPI model
type RegistrationAPI struct {
	Fullname        *string `json:"fullname,omitempty" example:"John doe" validate:"required"`
	Email           *string `json:"email,omitempty" example:"john.doe@mail.com" validate:"required,email"`
	Password        *string `json:"password,omitempty" example:"@Password123" validate:"required,min=8,regex=^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$"`
	ConfirmPassword *string `json:"confirm_password,omitempty" example:"@Password123" validate:"required,eqfield=Password"`
	RefferalCode    *string `json:"refferal_code,omitempty" example:"E6So5no5" validate:"required"`
}

// VerificationAccountAPI model
type VerificationAccountAPI struct {
	Email            *string `json:"email,omitempty" example:"john.doe@mail.com" validate:"required"`
	VerificationCode *string `json:"verification_code,omitempty" example:"7085" validate:"required"`
}

// LoginResponse model
type LoginResponse struct {
	Token    *ResponseToken `json:"token,omitempty"`
	Business *Business      `json:"business,omitempty"`
	User     *User          `json:"user,omitempty"`
}

// RegistrationResponse model
type RegistrationResponse struct {
	Business *Business `json:"business,omitempty"`
	User     *User     `json:"user,omitempty"`
}

// ResponseToken model
type ResponseToken struct {
	AccessToken  *string `json:"access_token,omitempty"`
	ExpiresIn    *int    `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	TokenType    *string `json:"token_type,omitempty"`
}

// ResponseAuthenticate model
type ResponseAuthenticate struct {
	ClientID            *string    `json:"ClientID,omitempty"`
	UserID              *string    `json:"UserID,omitempty"`
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
