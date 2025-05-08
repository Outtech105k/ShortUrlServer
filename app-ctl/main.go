package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	redisclient "github.com/Outtech105k/ShortUrlServer/app-ctl/redis-client"
	"github.com/Outtech105k/ShortUrlServer/app-ctl/routes"
	"github.com/Outtech105k/ShortUrlServer/app-ctl/utils"
)

func run() error {
	// Connct Redis
	redisAdapter, err := redisclient.NewRedisAdapter()
	if err != nil {
		return fmt.Errorf("connectRedis: %w", err)
	}
	defer redisAdapter.Close()

	// Setup AppContext
	appCtx := &utils.AppContext{
		Redis: *redisAdapter,
	}

	// Setup Gin Router
	router := routes.SetupRouter(appCtx)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	serverErrChan := make(chan error, 1)

	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrChan <- fmt.Errorf("server listen error: %w", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit: // 終了信号検知
		log.Printf("Received signal: %s. Initiating shutdown...\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("server shutdown failed: %w", err)
		}
		log.Println("Server gracefully stopped")
		return nil

	case err := <-serverErrChan: // ListenAndServe中のエラー検知
		return err
	}
}

func main() {
	if err := run(); err != nil {
		log.Println("Application exited with error:", err)
		os.Exit(1)
	}
}
