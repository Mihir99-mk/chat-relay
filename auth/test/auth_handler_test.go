package auth_test

import (
	"auth/config"
	"auth/config/mocks"
	"auth/internal/auth"
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService mocks the IService interface
type MockService struct {
	mock.Mock
}

func (m *MockService) SlackCallback(ctx context.Context, form auth.FormData) error {
	args := m.Called(ctx, form)
	return args.Error(0)
}

func setupHandler(_ *testing.T) (auth.IHandler, *MockService, *echo.Echo) {
	mockEnv := new(mocks.IEnv)
	mockEnv.On("SlackClientID").Return("client_id")
	mockEnv.On("SlackClientSecret").Return("client_secret")
	mockEnv.On("SlackRedirectURI").Return("http://localhost:8080/v1/api/auth/slack/callback")
	mockEnv.On("Get", mock.Anything).Return("dummy_value") // ✅ Important!

	mockCfg := new(mocks.IConfig)
	mockCfg.On("Env").Return(mockEnv)

	mockService := new(MockService)

	handler := auth.NewHandler(mockCfg)
	h := handler.(auth.HandlerAccessor).WithService(mockService)

	e := echo.New()
	return h, mockService, e
}

// Ensure handler exposes service injection for testing
type HandlerAccessor interface {
	auth.IHandler
	WithService(service auth.IService) auth.IHandler
}

func TestSlackLogin(t *testing.T) {
	envFile := "C:\\chatrelay\\auth\\.env"
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env file from %s: %s", envFile, err)
	}

	log.Println(".env file loaded successfully")
	// config.NewEnv()

	cfg := config.NewConfig()
	h := auth.NewHandler(cfg)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v1/api/auth/slack/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = h.SlackLogin(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
	assert.Contains(t, rec.Header().Get("Location"), "slack.com/oauth")
}

func TestSlackCallback_Success(t *testing.T) {
	h, mockService, e := setupHandler(t)

	form := "code=abc123&state=secure_random_state"
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/v1/api/auth/slack/callback", io.NopCloser(bytes.NewReader([]byte(form))))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("SlackCallback", mock.Anything, mock.Anything).Return(nil)

	err := h.SlackCallback(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Slack login successful")
}

func TestSlackCallback_Error(t *testing.T) {
	h, mockService, e := setupHandler(t)

	form := "code=fail&state=secure_random_state"
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/v1/api/auth/slack/callback", io.NopCloser(bytes.NewReader([]byte(form))))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("SlackCallback", mock.Anything, mock.Anything).Return(errors.New("invalid code"))

	err := h.SlackCallback(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "OAuth callback failed")
}
