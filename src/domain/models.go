package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Role      string         `json:"role" gorm:"not null"` // 'merchant' or 'customer'
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	MerchantID  uint           `json:"merchant_id" gorm:"not null"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Stock       int            `json:"stock" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Merchant    User           `json:"-" gorm:"foreignKey:MerchantID"`
}

type Transaction struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CustomerID   uint           `json:"customer_id" gorm:"not null"`
	ProductID    uint           `json:"product_id" gorm:"not null"`
	Quantity     int            `json:"quantity" gorm:"not null"`
	TotalPrice   float64        `json:"total_price" gorm:"not null"`
	Discount     float64        `json:"discount"`
	FreeShipping bool           `json:"free_shipping"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	Customer     User           `json:"-" gorm:"foreignKey:CustomerID"`
	Product      Product        `json:"-" gorm:"foreignKey:ProductID"`
}
