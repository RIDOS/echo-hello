package test

import (
	"bytes"
	"fmt"
	"github.com/RIDOS/echo-hello/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const roleAdmin = "admin"
const message = "red button user detected"

func TestHttpStatus(t *testing.T) {
	testCases := []struct {
		name          string
		role          string
		serverMessage string
	}{
		{
			name: "Test without role",
		},
		{
			name: "Test with role user",
			role: "User",
		},
		{
			name: "Test with role admin",
			role: roleAdmin,
		},
		{
			name: "Test with role admin in upper register",
			role: "ADMIN",
		},
	}

	s := service.New()

	dl := s.DaysLeft()
	str := fmt.Sprintf("Days left: %d", dl)

	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, str)
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.Logger = log.Output(&buf)

			req := httptest.NewRequest(http.MethodGet, "/status", nil)
			req.Header.Set("User-Role", testCase.role)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			roleCheckMiddleware := RoleCheck(handler)
			err := roleCheckMiddleware(c)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			if strings.EqualFold(testCase.role, roleAdmin) {
				assert.Contains(t, buf.String(), message)
			} else {
				assert.NotContains(t, buf.String(), message)
			}
		})
	}
}

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		usr := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(usr, roleAdmin) {
			log.Warn().Msg(message)
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
