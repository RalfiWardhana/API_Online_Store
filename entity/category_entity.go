package entity

import (
	"time"

	"gorm.io/gorm"
)

// 1 catgeory entity have many products
type Category struct {
	ID       string    `gorm:"primary_key" json:"id"`
	Name     string    `json:"name" gorm:"not null"`
	Products []Product `json:"-" gorm:"foreignkey:CategoryId"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
