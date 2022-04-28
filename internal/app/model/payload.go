package model

// PayloadToken a struct to store all payload for token
type PayloadToken struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}
