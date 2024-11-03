package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		usr := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(usr, roleAdmin) {
			log.Warn().Msg("red button user detected")
		}

		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}
