package app

import (
	"fmt"
	"github.com/RIDOS/echo-hello/internal/app/endpoint"
	"github.com/RIDOS/echo-hello/internal/app/mw"
	"github.com/RIDOS/echo-hello/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.GET("/status", a.e.Status, mw.RoleCheck)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running...")
	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}

	return nil
}
