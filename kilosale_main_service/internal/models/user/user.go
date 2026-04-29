package models

import (
	"time"

	"github.com/gechderib/kilosale/pkg/common"
)

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
	Name     *string `json:"name"`
	Phone    *string `json:"phone"`
	Password *string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (r UserRequest) validate() map[string]string {
	errors := make(map[string]string)

	if r.Name == nil {
		errors["name"] = "Name is required"
	} else if *r.Name == "" {
		errors["name"] = "Name Cannot be empty"
	}

	if r.Phone == nil {
		errors["phone"] = "Phone is required"
	} else if *r.Phone == "" {
		errors["phone"] = "Phone Cannot be empty"
	} else if len(*r.Phone) < 10 {
		errors["phone"] = "Phone must be at least 10 characters long"
	} else if len(*r.Phone) > 15 {
		errors["phone"] = "Phone must be at most 15 characters long"
	}

	if r.Password == nil {
		errors["password"] = "Password is required"
	} else if *r.Password == "" {
		errors["password"] = "Password Cannot be empty"
	} else if len(*r.Password) < 8 {
		errors["password"] = "Password must be at least 8 characters long"
	}

	return errors
}

func (r UserRequest) ToUser() *User {
	return &User{
		Name:     *r.Name,
		Phone:    *r.Phone,
		Password: *r.Password,
	}
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
