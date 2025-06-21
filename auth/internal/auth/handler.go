package auth

import (
	"auth/config"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	tracer  trace.Tracer
}

func NewHandler(cfg config.IConfig) IHandler {
	return handler{
		service: NewService(cfg),
		env:     cfg.Env(),
		tracer:  otel.Tracer("auth/handler"),
	}
}

func (h handler) SlackCallback(c echo.Context) error {
	ctx, span := h.tracer.Start(c.Request().Context(), "SlackCallback")
	defer span.End()

	code := c.QueryParam("code")
	span.SetAttributes(attribute.String("http.method", c.Request().Method))
	span.SetAttributes(attribute.String("http.route", c.Path()))
	span.SetAttributes(attribute.String("slack.oauth.code", code))

	if code == "" {
		span.RecordError(fmt.Errorf("missing code parameter"))
		span.SetStatus(1, "missing 'code' parameter in Slack callback")
		return c.String(http.StatusBadRequest, "Missing code parameter")
	}

	if err := h.service.SlackCallback(ctx, code); err != nil {
		span.RecordError(err)
		span.SetStatus(1, "SlackCallback service error")
		return err
	}

	span.SetStatus(0, "Slack OAuth callback successful")
	return c.JSON(http.StatusOK, "Data fetched!!!")
}

func (h handler) SlackLogin(c echo.Context) error {
	_, span := h.tracer.Start(c.Request().Context(), "SlackLogin")
	defer span.End()

	scopes := "app_mentions:read,chat:write,channels:history,groups:history,im:history"

	authURL := fmt.Sprintf(
		"https://slack.com/oauth/v2/authorize?client_id=%s&scope=%s&redirect_uri=%s",
		h.env.GetSlackClientId(), scopes, url.QueryEscape(h.env.GetSlackRedirectUrl()),
	)

	span.SetAttributes(
		attribute.String("http.method", c.Request().Method),
		attribute.String("http.route", c.Path()),
		attribute.String("slack.auth_url", authURL),
	)

	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}
