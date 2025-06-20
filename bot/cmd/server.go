package cmd

import (
	"bot/config"
	"bot/route"
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
)

func StartServer() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := config.Migrate(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	shutdowns := config.InitOtel(ctx)
	defer func() {
		for _, shutdown := range shutdowns {
			if err := shutdown(ctx); err != nil {
				log.Printf("OTel shutdown error: %v", err)
			}
		}
	}()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = echox.HttpErrorHandler

	route.InitRoute(e)

	go func() {
		env := config.NewEnv()
		port := fmt.Sprintf(":%v", env.GetPort())
		if err := e.Start(port); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}

		log.Printf("Server started :%v ", env.GetPort())
	}()

	<-ctx.Done()
	log.Println("shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server shut down cleanly")
}
