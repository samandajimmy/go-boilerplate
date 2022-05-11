package model

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Token to store JWT token data
type Token struct {
	Name   string `json:"name"`
	Claims jwt.RegisteredClaims
}

// AccountToken to store account token data
type AccountToken struct {
	ID        int64         `json:"id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	Token     string        `json:"token"`
	Status    string        `json:"status"`
	ExpiredAt sql.NullTime  `json:"expiredAt" db:"expired_at"`
	CreatedAt time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time     `json:"updatedAt" db:"updated_at"`
	ExpiresIn time.Duration `json:"expiresIn" db:"-"`
}
