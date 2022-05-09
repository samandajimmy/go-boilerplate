package http

import (
	cmdutil "go-boiler-plate/cmd/util"
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/payload"

	"github.com/labstack/echo/v4"

	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

type TokenHandler struct {
	Ihandler      pgdutil.IHandler
	ITokenUsecase token.ITokenUsecase
}

// NewTokensHandler represent to register tokens endpoint
func NewTokensHandler(echoGroup cmdutil.EchoGroup, iTknUsecase token.ITokenUsecase) {
	handler := &TokenHandler{
		Ihandler:      pgdutil.NewHandler(&pgdutil.Handler{}),
		ITokenUsecase: iTknUsecase,
	}

	echoGroup.Token.POST("/create", handler.HCreateToken)
	echoGroup.Token.GET("/get", handler.HGetToken)
	echoGroup.Token.GET("/refresh", handler.HRefreshToken)
}

func (th *TokenHandler) HCreateToken(c echo.Context) error {
	var pl payload.TokenRequest
	var errors pgdutil.ResponseErrors
	err := th.Ihandler.Validate(c, &pl)

	if err != nil {
		return th.Ihandler.ShowResponse(c, nil, err, errors)
	}

	response, err := th.ITokenUsecase.UCreateToken(c, pl)

	return th.Ihandler.ShowResponse(c, response, err, errors)
}

func (th *TokenHandler) HGetToken(c echo.Context) error {
	var pl payload.TokenRequest
	var errors pgdutil.ResponseErrors

	_ = echo.QueryParamsBinder(c).
		String("username", &pl.Username).
		String("password", &pl.Password).
		BindError()

	err := th.Ihandler.Validate(c, &pl)

	if err != nil {
		return th.Ihandler.ShowResponse(c, nil, err, errors)
	}

	accToken, err := th.ITokenUsecase.UGetToken(c, pl.Username, pl.Password)

	return th.Ihandler.ShowResponse(c, accToken, err, errors)
}

func (th *TokenHandler) HRefreshToken(c echo.Context) error {
	var pl payload.TokenRequest
	var errors pgdutil.ResponseErrors
	_ = echo.QueryParamsBinder(c).
		String("username", &pl.Username).
		String("password", &pl.Password).
		BindError()

	err := th.Ihandler.Validate(c, &pl)

	if err != nil {
		return th.Ihandler.ShowResponse(c, nil, err, errors)
	}

	accToken, err := th.ITokenUsecase.URefreshToken(c, pl.Username, pl.Password)

	return th.Ihandler.ShowResponse(c, accToken, err, errors)
}
