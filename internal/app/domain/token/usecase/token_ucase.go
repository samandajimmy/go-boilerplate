package usecase

import (
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/app/payload"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"repo.pegadaian.co.id/ms-pds/modules/pgdutil"
)

type tokenUseCase struct {
	tokenRepo token.ITokenRepository
}

// NewTokenUsecase will create new an TokenUseCase object representation of Tokens.UseCase interface
func NewTokenUsecase(tkn token.ITokenRepository) token.ITokenUsecase {
	return &tokenUseCase{
		tokenRepo: tkn,
	}
}

func (tkn *tokenUseCase) UCreateToken(c echo.Context, accToken model.AccountToken) (payload.TokenResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(accToken.Password), bcrypt.DefaultCost)
	accToken.Password = string(hashedPassword)
	err := tkn.tokenRepo.RCreate(c, &accToken)

	if err != nil {
		return payload.TokenResponse{}, model.ErrCreateToken
	}

	return payload.TokenResponse{
		Username:  accToken.Username,
		Token:     accToken.Token,
		Status:    accToken.Status,
		CreatedAt: accToken.CreatedAt,
		UpdatedAt: accToken.UpdatedAt,
		ExpiresIn: accToken.ExpiresIn,
	}, nil
}

func (tkn *tokenUseCase) UGetToken(c echo.Context, username string, password string) (payload.TokenResponse, error) {
	accToken := &model.AccountToken{}
	accToken.Username = username

	// get account
	err := tkn.tokenRepo.RGetByUsername(c, accToken)

	if err != nil {
		return payload.TokenResponse{}, model.ErrUsername
	}

	if err = verifyToken(accToken, password, false); err != nil {
		return payload.TokenResponse{}, err
	}

	return payload.TokenResponse{
		Username:  accToken.Username,
		Token:     accToken.Token,
		Status:    accToken.Status,
		CreatedAt: accToken.CreatedAt,
		UpdatedAt: accToken.UpdatedAt,
		ExpiresIn: accToken.ExpiresIn,
	}, nil
}

func (tkn *tokenUseCase) URefreshToken(c echo.Context, username string, password string) (payload.TokenResponse, error) {
	accToken := &model.AccountToken{}
	accToken.Username = username

	// get account
	err := tkn.tokenRepo.RGetByUsername(c, accToken)

	if err != nil {
		return payload.TokenResponse{}, model.ErrUsername
	}

	if err = verifyToken(accToken, password, true); err != nil {
		return payload.TokenResponse{}, err
	}

	// refresh JWT
	err = tkn.tokenRepo.RUpdateToken(c, accToken)

	if err != nil {
		return payload.TokenResponse{}, model.ErrCreateToken
	}

	err = tkn.tokenRepo.RGetByUsername(c, accToken)

	if err != nil {
		return payload.TokenResponse{}, err
	}

	return payload.TokenResponse{
		Username:  accToken.Username,
		Token:     accToken.Token,
		Status:    accToken.Status,
		CreatedAt: accToken.CreatedAt,
		UpdatedAt: accToken.UpdatedAt,
		ExpiresIn: accToken.ExpiresIn,
	}, nil
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
	if accToken.ExpiredAt.Time.Before(now) && !isUpdate {
		return model.ErrTokenExpired
	}

	return nil
}
