package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/source-c0de/contacthub/internal/config"
	"github.com/source-c0de/contacthub/internal/handler"
	"github.com/source-c0de/contacthub/internal/middleware"
	"github.com/source-c0de/contacthub/internal/repository"
	"github.com/source-c0de/contacthub/internal/service"

	"go.uber.org/zap"
)

func Setup(
	cfg *config.Config,
	logger *zap.Logger,
	db *sqlx.DB,
) *gin.Engine {

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Logger
	r.Use(middleware.Logger(logger))
	r.Use(gin.Recovery())

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	authSvc := service.NewAuthService(cfg, userRepo)

	// Initialize handlers
	authHandler := &handler.AuthHandler{AuthSvc: authSvc}

	// Auth Routes (public)
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/logout", authHandler.Logout)
	}

	// Protected Routes (require JWT)
	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth(cfg, authSvc))
	{
		api.GET("/me", authHandler.Me)
	}

	return r
}
