package main

import (
	"boilerplate/config"
	"boilerplate/internal/routes"
	"boilerplate/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// Inicializa configurações.
	cfg := config.Load()

	// Inicializa o logger.
	log := logger.Init(cfg.LogLevel)

	// Configura as rotas do Gin.
	router := routes.InitRouter(log)

	// Inicia o servidor.
	log.Info("Starting server...", zap.String("port", cfg.Port))
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Server failed to start", zap.Error(err))
	}
}
