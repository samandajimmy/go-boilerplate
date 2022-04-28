package model

import "github.com/labstack/echo"

// EchoGroup to store routes group
type EchoGroup struct {
	Token *echo.Group
}
