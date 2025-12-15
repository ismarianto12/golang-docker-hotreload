package entities

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(200);not null"`
	Description string    `json:"description" gorm:"type:text";null`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int       `json:"stock" gorm:"not null;default:0";null`
	UserID      uint      `json:"user_id" gorm:"not null;index"`
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID";null`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime";null`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime";null`
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}
