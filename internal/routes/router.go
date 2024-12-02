package routes

import (
	"boilerplate/internal/handlers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitRouter inicializa as rotas da aplicação.
func InitRouter(log *zap.Logger) *gin.Engine {
	r := gin.New()

	// Middlewares globais.
	r.Use(gin.Recovery())
	r.Use(loggingMiddleware(log))

	// Grupo API v1.
	apiV1 := r.Group("/api/v1")
	{
		registerHealthRoutes(apiV1)
		// Adicione outros grupos de recursos aqui.
	}

	return r
}

// loggingMiddleware insere o logger no contexto.
func loggingMiddleware(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", log)
		c.Next()
	}
}

// registerHealthRoutes registra as rotas relacionadas à saúde do sistema.
func registerHealthRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/health")
	{
		health.GET("", handlers.Health)
		// Adicione outros endpoints relacionados a 'health', se necessário.
	}
}
