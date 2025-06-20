package auth

import (
	"auth/config"
	"auth/internal/auth"
	"auth/util"

	"github.com/labstack/echo/v4"
)

func InitSlackRoute(e *echo.Group) {
	authGroup := e.Group(util.Auth)
	config := config.NewConfig()
	authHandler := auth.NewHandler(config)

	slackAuth := authGroup.Group(util.Slack)
	slackAuth.GET("/login", authHandler.SlackLogin)
	slackAuth.GET("/callback", authHandler.SlackCallback)
}
