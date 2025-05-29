package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

type User struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn  time.Time `json:"voided_on" gorm:"default:NULL"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role" gorm:"not null"`
	Csrf string `json:"csrf" gorm:"not null"`
}
