package main

import (
	"database/sql"
	"go-boiler-plate/internal/app/domain/token"
	_tokenRepository "go-boiler-plate/internal/app/domain/token/repository"

	"github.com/go-pg/pg/v9"
)

type Repositories struct {
	TokenRepository token.Repository
}

func newRepositories(dbConn *sql.DB, dbBun *pg.DB) Repositories {

	tokenRepository := _tokenRepository.NewPsqlTokenRepository(dbConn)

	return Repositories{
		TokenRepository: tokenRepository,
	}

}
