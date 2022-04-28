package main

import (
	"go-boiler-plate/internal/app/domain/token"
	_tokenUseCase "go-boiler-plate/internal/app/domain/token/usecase"
	"time"
)

type Usecases struct {
	TokenUseCase token.UseCase
}

func newUsecases(repos Repositories, timeout time.Duration) Usecases {

	tokenUseCase := _tokenUseCase.NewTokenUseCase(repos.TokenRepository, timeout)

	return Usecases{TokenUseCase: tokenUseCase}
}
