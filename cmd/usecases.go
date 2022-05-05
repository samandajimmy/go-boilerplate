package main

import (
	_tokenUsecase "go-boiler-plate/internal/app/domain/token/usecase"

	"go-boiler-plate/internal/app/domain/token"
)

type Usecases struct {
	ITokenUseCase token.ITokenUsecase
}

func newUsecases(repos Repositories) Usecases {

	tokenUseCase := _tokenUsecase.NewTokenUsecase(repos.ITokenRepository)

	return Usecases{ITokenUseCase: tokenUseCase}
}
