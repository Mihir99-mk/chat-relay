package auth

import (
	"auth/config"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type IHandler interface {
	SlackLogin(c echo.Context) error
	SlackCallback(c echo.Context) error
}

type Handler interface {
	IHandler
}

type handler struct {
	service IService
	env     config.IEnv
	logger  *otelzap.Logger
}

func NewHandler(cfg config.IConfig) IHandler {
	return handler{
		service: NewService(cfg),
		env:     cfg.Env(),
		logger:  config.Logger(),
	}
}

func getRequestFields(c echo.Context) []zap.Field {
	return []zap.Field{
		zap.String("trace_id", trace.SpanContextFromContext(c.Request().Context()).TraceID().String()),
		zap.String("method", c.Request().Method),
		zap.String("path", c.Path()),
		zap.String("remote_ip", c.RealIP()),
		zap.String("user_agent", c.Request().UserAgent()),
	}
}

func (h handler) SlackCallback(c echo.Context) error {
	ctx := c.Request().Context()
	log := h.logger.With(getRequestFields(c)...)
	code := c.QueryParam("code")

	if code == "" {
		log.Error("Missing 'code' parameter in Slack callback")
		return c.String(http.StatusBadRequest, "Missing code parameter")
	}

	log.Info("Processing Slack OAuth callback", zap.String("code", code))

	if err := h.service.SlackCallback(ctx, code); err != nil {
		log.Error("SlackCallback service error", zap.Error(err))
		return err
	}

	log.Info("Slack OAuth callback successful")
	return c.JSON(http.StatusOK, "Data fetched!!!")
}

func (h handler) SlackLogin(c echo.Context) error {
	log := h.logger.With(getRequestFields(c)...)
	scopes := "app_mentions:read,chat:write,channels:history,groups:history,im:history"

	authURL := fmt.Sprintf(
		"https://slack.com/oauth/v2/authorize?client_id=%s&scope=%s&redirect_uri=%s",
		h.env.GetSlackClientId(), scopes, url.QueryEscape(h.env.GetSlackRedirectUrl()),
	)

	log.Info("Redirecting to Slack OAuth login",
		append(getRequestFields(c), zap.String("auth_url", authURL))...,
	)

	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}
