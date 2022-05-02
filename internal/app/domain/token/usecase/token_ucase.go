package usecase

import (
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/model"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

type tokenUseCase struct {
	tokenRepo      token.Repository
	contextTimeout time.Duration
}

// NewTokenUseCase will create new an TokenUseCase object representation of Tokens.UseCase interface
func NewTokenUseCase(tkn token.Repository, timeout time.Duration) token.UseCase {
	return &tokenUseCase{
		tokenRepo:      tkn,
		contextTimeout: timeout,
	}
}

func (tkn *tokenUseCase) UCreateToken(c echo.Context, accToken *model.AccountToken) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(accToken.Password), bcrypt.DefaultCost)
	accToken.Password = string(hashedPassword)
	err := tkn.tokenRepo.RCreate(c, accToken)

	if err != nil {
		return model.ErrCreateToken
	}

	return nil
}

func (tkn *tokenUseCase) UGetToken(c echo.Context, username string, password string) (*model.AccountToken, error) {
	accToken := &model.AccountToken{}
	accToken.Username = username

	// get account
	err := tkn.tokenRepo.RGetByUsername(c, accToken)

	if err != nil {
		return nil, model.ErrUsername
	}

	if err = verifyToken(accToken, password, false); err != nil {
		return nil, err
	}

	// rearrange accountToken
	accToken.ID = 0
	accToken.Password = ""
	accToken.Status = nil

	return accToken, nil
}

func (tkn *tokenUseCase) URefreshToken(c echo.Context, username string, password string) (*model.AccountToken, error) {
	accToken := &model.AccountToken{}
	accToken.Username = username

	// get account
	err := tkn.tokenRepo.RGetByUsername(c, accToken)

	if err != nil {
		return nil, model.ErrUsername
	}

	if err = verifyToken(accToken, password, true); err != nil {
		return nil, err
	}

	// refresh JWT
	err = tkn.tokenRepo.RUpdateToken(c, accToken)

	if err != nil {
		return nil, model.ErrCreateToken
	}

	_ = tkn.tokenRepo.RGetByUsername(c, accToken)

	// rearrange accountToken
	accToken.ID = 0
	accToken.Password = ""
	accToken.Status = nil

	return accToken, nil
}

func (tkn *tokenUseCase) URefreshAllToken() error {
	// update all account token data
	err := tkn.tokenRepo.RUpdateAllAccountTokenExpiry()

	if err != nil {
		return err
	}

	return nil
}

func verifyToken(accToken *model.AccountToken, password string, isUpdate bool) error {
	now := pgdutil.NowUTC()
	// validate account
	// check password
	err := bcrypt.CompareHashAndPassword([]byte(accToken.Password), []byte(password))

	if err != nil {
		return model.ErrPassword
	}

	// token availabilty
	if accToken.ExpireAt.Before(now) && !isUpdate {
		return model.ErrTokenExpired
	}

	return nil
}
