package route

import (
	"bot/config"
	"bot/internal/healthcheck"
	"bot/route/event"
	"bot/util"

	mock "bot/internal/slack"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	v1Group := e.Group(util.V1)

	config := config.NewConfig()
	v1Group.GET(util.Healthz, healthcheck.HealthCheckHandler(config))

	v1Group = v1Group.Group(util.Api)

	// Mock bot streaming API
	handler := mock.NewMockHandler()

	v1Group.POST("/chat/stream", handler.StreamHandler)

	// bot event
	go event.InitSlackEvent()
}
