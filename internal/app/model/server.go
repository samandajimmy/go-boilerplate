package model

import "github.com/labstack/echo/v4"

// EchoGroup to store routes group
type EchoGroup struct {
	Api   *echo.Group
	Token *echo.Group
}
