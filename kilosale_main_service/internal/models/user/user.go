package models

import "github.com/gechderib/kilosale/pkg/common"

// For Database Model
type User struct {
	common.UUID
	Name     string `gorm:"not null;" json:"name"`
	Phone    string `gorm:"unique;not null" json:"phone"`
	Password string `gorm:"not null;size:255;max:255;min:8" json:"-"`
	common.CustomTimeStamp
}

// For RequestResponse
type UserRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
