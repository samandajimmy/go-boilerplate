package repository

import (
	"go-boiler-plate/internal/app/domain/token"
	"go-boiler-plate/internal/app/model"
	"go-boiler-plate/internal/pkg/database"
	"go-boiler-plate/internal/pkg/msg"

	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/samandajimmy/pgdlogger"
	"github.com/samandajimmy/pgdutil"
)

type psqlTokenRepository struct {
	db *database.Db
}

// NewPsqlTokenRepository will create an object that represent the token.Repository interface
func NewPsqlTokenRepository(db *database.Db) token.ITokenRepository {
	return &psqlTokenRepository{db}
}

func (tknRepo *psqlTokenRepository) RCreate(c echo.Context, accToken *model.AccountToken) error {
	var lastID int64
	now := time.Now()
	tokenExp := now.Add(stringToDuration(os.Getenv(`APP_JWT_TOKEN_EXP`)) * time.Second)
	token, err := createJWTToken(accToken, now, tokenExp)

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	accToken.Token = token
	accToken.ExpiredAt.Time = tokenExp
	accToken.CreatedAt = now
	accToken.UpdatedAt = now
	accToken.Status = msg.DefStatusActive

	query := `INSERT INTO account_tokens (username, password, token, expired_at, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = tknRepo.db.Sqlx.QueryRowx(query, accToken.Username, accToken.Password, accToken.Token,
		accToken.ExpiredAt.Time, accToken.Status, accToken.CreatedAt, accToken.UpdatedAt).Scan(&lastID)

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	accToken.ID = lastID
	accToken.ExpiresIn = time.Duration(time.Until(accToken.ExpiredAt.Time).Seconds())

	return nil
}

func (tknRepo *psqlTokenRepository) RGetByUsername(c echo.Context, accToken *model.AccountToken) error {
	query := `SELECT id, username, password, token, expired_at, status, updated_at, created_at
		FROM account_tokens WHERE status = $1 AND username = $2 limit 1;`

	err := tknRepo.db.Sqlx.QueryRowx(query, msg.DefStatusActive, accToken.Username).StructScan(accToken)

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	accToken.ExpiresIn = time.Duration(accToken.ExpiredAt.Time.Sub(pgdutil.NowUTC()).Seconds())

	return nil
}

func (tknRepo *psqlTokenRepository) RUpdateToken(c echo.Context, accToken *model.AccountToken) error {
	var id int64
	now := time.Now()
	tokenExp := now.Add(stringToDuration(os.Getenv(`APP_JWT_TOKEN_EXP`)) * time.Second)
	token, err := createJWTToken(accToken, now, tokenExp)

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	query := `UPDATE account_tokens SET token = $1, expired_at = $2, updated_at = $3 WHERE username = $4 RETURNING id`
	err = tknRepo.db.Sqlx.QueryRowx(query, token, tokenExp, now, accToken.Username).Scan(&id)

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	return nil
}

func (tknRepo *psqlTokenRepository) RUpdateAllAccountTokenExpiry() error {
	query := `UPDATE account_tokens SET expired_at = $1, updated_at = $2`
	_, err := tknRepo.db.Sqlx.Exec(query, nil, time.Now())

	if err != nil {
		pgdlogger.Make().Debug(err)

		return err
	}

	return nil
}

func stringToDuration(str string) time.Duration {
	hours, _ := strconv.Atoi(str)

	return time.Duration(hours)
}

func createJWTToken(accountToken *model.AccountToken, now time.Time, tokenExp time.Time) (string, error) {
	// nolint
	jwtDate := &jwt.NumericDate{tokenExp}
	token := model.Token{
		Name:   accountToken.Username,
		Claims: jwt.RegisteredClaims{ID: accountToken.Username, ExpiresAt: jwtDate},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, token.Claims)
	return rawToken.SignedString([]byte(os.Getenv(`APP_JWT_SECRET`)))
}
