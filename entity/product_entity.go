package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Stock       int    `json:"stock" gorm:"type:int;not null"`
	Price       int    `json:"price" gorm:"type:int;not null"`
	Image       string `json:"image" gorm:"type:varchar(255);not null"`
	UserId      string `json:"user_id" gorm:"type:varchar(255);not null"`
	CategoryId  string `json:"category_id" gorm:"type:varchar(255);not null"`

	User User `json:"-" gorm:"foreignkey:UserId"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
