package slack

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IMockHandler interface {
	StreamHandler(c echo.Context) error
}

type mockHandler struct {
}

func NewMockHandler() IMockHandler {
	return &mockHandler{}
}

func (m *mockHandler) StreamHandler(c echo.Context) error {
	var req ChatRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	chunks := []string{
		"Goroutines are lightweight...",
		"They allow for massive concurrency.",
		"Used heavily in Go's concurrency model.",
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())
	for _, chunk := range chunks {
		time.Sleep(1 * time.Second)
		_ = enc.Encode(StreamResponse{TextChunk: chunk})
		c.Response().Flush()
	}

	return enc.Encode(StreamResponse{Status: "done"})
}
