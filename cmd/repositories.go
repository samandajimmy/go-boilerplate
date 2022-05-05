package main

import (
	_tokenRepository "go-boiler-plate/internal/app/domain/token/repository"

	"go-boiler-plate/internal/app/domain/token"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ITokenRepository token.ITokenRepository
}

func newRepositories(sqlx *sqlx.DB) Repositories {

	tokenRepository := _tokenRepository.NewPsqlTokenRepository(sqlx)

	return Repositories{
		ITokenRepository: tokenRepository,
	}

}
