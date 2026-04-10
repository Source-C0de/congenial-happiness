package main

import (
	"fmt"
	"log"

	"github.com/source-c0de/contacthub/internal/config"
	"github.com/source-c0de/contacthub/internal/database"
	"github.com/source-c0de/contacthub/internal/handler"
	"github.com/source-c0de/contacthub/internal/router"
	"github.com/source-c0de/contacthub/internal/server"
	"github.com/source-c0de/contacthub/internal/service"
	"go.uber.org/zap"
)

var (
	Version   string
	BuildTime string
	GitCommit string
)

func main() {

	// 1. Load Config
	cfg := config.Load()

	// 2. Initialize Logger
	logger, err := zap.NewProduction()
	if cfg.Environment == "development" {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// 3. Database Connection
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// 4. Initialize Repositories

	// 5. Initialize Services
	var authSvc service.AuthService // Placeholder

	// 6. Initialize Handlers
	authHandler := &handler.AuthHandler{AuthSvc: authSvc}

	// 7.Router
	r := router.Setup(cfg, logger, authSvc, authHandler)

	// 8. Start Server
	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Info("Server Starting", zap.String("addr", addr), zap.String("env", cfg.Environment))

	srv := server.New(cfg, logger, r)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
