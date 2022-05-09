package msg

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError to store internal server error message
	ErrInternalServerError = errors.New("internal server error")

	// ErrMigrateNoChange to store migration no change error message
	ErrMigrateNoChange = errors.New("no change")

	// ErrEnvFileNF to store error env file not found error message
	ErrEnvFileNF = errors.New("env file not found! expected we have set system variable")

	// ErrCreateToken to create/update token error message
	ErrCreateToken = errors.New("terjadi kesalaahan saat membuat token")

	// ErrUsername to store username error message
	ErrUsername = errors.New("username atau password yang digunakan tidak valid")

	// ErrPassword to store password error message
	ErrPassword = errors.New("username atau password yang digunakan tidak valid")

	// ErrTokenExpired to store password error message
	ErrTokenExpired = errors.New("token anda telah kadarluarsa")
)

// DynamicErr to return parameterize errors
func DynamicErr(message string, args []interface{}) error {
	return fmt.Errorf(message, args...)
}
