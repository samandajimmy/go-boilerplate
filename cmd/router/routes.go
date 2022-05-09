package router

import (
	cmdutil "go-boiler-plate/cmd/util"
	_tokenHttpDelivery "go-boiler-plate/internal/app/domain/token/delivery/http"
	"go-boiler-plate/internal/pkg/database"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo         *echo.Echo
	EchoGroup    cmdutil.EchoGroup
	Repositories repositories
	Usecases     usecases
}

func NewRoutes(db *database.Db) Router {
	e := echo.New()
	e.Debug = false

	eGroup := cmdutil.EchoGroup{
		Api:   e.Group("/api"),
		Token: e.Group("/token"),
	}

	repos := newRepositories(db)

	return Router{
		Echo:         e,
		EchoGroup:    eGroup,
		Repositories: repos,
		Usecases:     newUsecases(repos),
	}
}

func (r *Router) LoadHandlers() {
	_tokenHttpDelivery.NewTokensHandler(r.EchoGroup, r.Usecases.ITokenUsecase)
}
