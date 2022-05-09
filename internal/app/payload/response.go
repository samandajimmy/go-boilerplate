package payload

import (
	"time"
)

type TokenResponse struct {
	Username  string        `json:"username"`
	Token     string        `json:"token"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	ExpiresIn time.Duration `json:"expiresIn"`
}
