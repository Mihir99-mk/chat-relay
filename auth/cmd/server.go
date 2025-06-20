package cmd

import (
	"auth/config"
	"auth/route"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mihir99-mk/chat-relay-lib/echox"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func StartServer() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Run database migrations
	if err := config.Migrate(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// Initialize observability (tracing, metrics, logging, etc.)
	shutdowns := config.InitOtel(ctx)
	defer func() {
		for _, shutdown := range shutdowns {
			if err := shutdown(ctx); err != nil {
				log.Printf("OTel shutdown error: %v", err)
			}
		}
	}()

	// Set up Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(otelecho.Middleware(config.NewEnv().GetServiceName()))

	e.HTTPErrorHandler = echox.HttpErrorHandler

	// Register routes here
	route.InitRoute(e)

	// Start server in a goroutine to support graceful shutdown
	go func() {
		env := config.NewEnv()
		port := fmt.Sprintf(":%v", env.GetPort())
		if err := e.Start(port); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	log.Println("shutting down server gracefully...")

	// Context with timeout for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server shut down cleanly")
}
