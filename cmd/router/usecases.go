package router

import (
	_tokenUsecase "go-boiler-plate/internal/app/domain/token/usecase"

	"go-boiler-plate/internal/app/domain/token"
)

type usecases struct {
	ITokenUsecase token.ITokenUsecase
}

func newUsecases(repos repositories) usecases {
	tokenUsecase := _tokenUsecase.NewTokenUsecase(repos.ITokenRepository)

	return usecases{ITokenUsecase: tokenUsecase}
}
