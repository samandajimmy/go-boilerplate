package router

import (
	_tokenRepository "go-boiler-plate/internal/app/domain/token/repository"

	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/pkg/database"
)

type repositories struct {
	ITokenRepository token.ITokenRepository
}

func newRepositories(db *database.Db) repositories {
	tokenRepository := _tokenRepository.NewPsqlTokenRepository(db)

	return repositories{
		ITokenRepository: tokenRepository,
	}

}
