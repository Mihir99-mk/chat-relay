package route

import (
	"auth/config"
	"auth/internal/healthcheck"
	"auth/route/auth"
	"auth/util"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	v1Group := e.Group(util.V1)

	config := config.NewConfig()
	v1Group.GET(util.Healthz, healthcheck.HealthCheckHandler(config))

	v1Group = v1Group.Group(util.Api)

	auth.InitSlackRoute(v1Group)
}
