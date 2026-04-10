package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/source-c0de/contacthub/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	logger *zap.Logger
	router *gin.Engine
}

func New(cfg *config.Config, logger *zap.Logger, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}
}
