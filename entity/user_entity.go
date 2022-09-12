package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         string `json:"id" gorm:"primary_key;not null"`
	Name       string `json:"name" gorm:"not null" validate:"required"`
	Email      string `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password   string `json:"password" gorm:"not null" validate:"required,min=8,max=100"`
	Otp        string `json:"otp" gorm:"not null"`
	OtpExpired int64  `json:"otp_expiry" gorm:"not null"`
	RoleId     int    `json:"role_id" gorm:"not null;default:1"`

	Role Role `json:"-" gorm:"foreignkey:RoleId"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
