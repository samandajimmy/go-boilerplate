package delivery

import (
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/model"

	"github.com/labstack/echo"
)

type TokensHandler struct {
	response     model.Response
	respErrors   model.ResponseErrors
	TokenUseCase token.UseCase
}

// NewTokensHandler represent to register tokens endpoint
func NewTokensHandler(echoGroup model.EchoGroup, tknUseCase token.UseCase) {
	handler := &TokensHandler{
		TokenUseCase: tknUseCase,
	}

	echoGroup.Token.POST("/create", handler.HCreateToken)
	echoGroup.Token.GET("/get", handler.HGetToken)
	echoGroup.Token.GET("/refresh", handler.HRefreshToken)
}

func (tkn *TokensHandler) HCreateToken(c echo.Context) error {
	var accountToken model.AccountToken
	tkn.response, tkn.respErrors = model.NewResponse()
	err := c.Bind(&accountToken)

	if err != nil {
		tkn.respErrors.SetTitle(model.MessageUnprocessableEntity)
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	err = tkn.TokenUseCase.UCreateToken(c, &accountToken)

	if err != nil {
		tkn.respErrors.SetTitle(err.Error())
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	tkn.response.SetResponse(accountToken, &tkn.respErrors)

	return tkn.response.Body(c, err)
}

func (tkn *TokensHandler) HGetToken(c echo.Context) error {
	tkn.response, tkn.respErrors = model.NewResponse()
	var getToken model.PayloadToken

	if err := c.Bind(&getToken); err != nil {
		tkn.respErrors.SetTitle(model.MessageUnprocessableEntity)
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	accToken, err := tkn.TokenUseCase.UGetToken(c, getToken.UserName, getToken.Password)

	if err != nil {
		tkn.respErrors.SetTitle(err.Error())
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	tkn.response.SetResponse(accToken, &tkn.respErrors)
	return tkn.response.Body(c, err)
}

func (tkn *TokensHandler) HRefreshToken(c echo.Context) error {
	tkn.response, tkn.respErrors = model.NewResponse()
	var refToken model.PayloadToken
	if err := c.Bind(&refToken); err != nil {
		tkn.respErrors.SetTitle(model.MessageUnprocessableEntity)
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	accToken, err := tkn.TokenUseCase.URefreshToken(c, refToken.UserName, refToken.Password)

	if err != nil {
		tkn.respErrors.SetTitle(err.Error())
		tkn.response.SetResponse("", &tkn.respErrors)

		return tkn.response.Body(c, err)
	}

	tkn.response.SetResponse(accToken, &tkn.respErrors)

	return tkn.response.Body(c, err)
}
