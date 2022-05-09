package util

import (
	"golang.org/x/crypto/bcrypt"
	"repo.pegadaian.co.id/ms-pds/modules/pgdlogger"
)

type IUtil interface {
	BcryptHashedPassword(password string) string
}

type UtilSvc struct{}

func NewUtil() IUtil {
	return &UtilSvc{}
}

func (utl *UtilSvc) BcryptHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		pgdlogger.Make().Panic(err)
	}

	return string(hashedPassword)
}
