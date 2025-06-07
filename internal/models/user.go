package models

import (
	"time"
)

type User struct {
	Id         uint      `json:"id" gorm:"primary_key" unique:"true"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email" gorm:"index;unique;not null"`
	Password   string    `json:"password" gorm:"not null"`
	Role       string    `json:"role" gorm:"not null"`
	IsVerified bool      `json:"isVerified" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn   time.Time `json:"voided_on" gorm:"default:NULL"`
}
