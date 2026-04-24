package models

import "github.com/gechderib/kilosale/pkg/common"

type User struct {
	common.BaseModel
	common.CustomTimeStamp
	Name     string `gorm:"not null" json:"name"`
	Phone    string `gorm:"unique;not null" json:"phone"`
	Password string `gorm:"not null;size:255;max:255;min:8" json:"-"`
}
