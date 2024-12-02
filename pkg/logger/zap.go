package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init configura o logger Zap com suporte a logging em arquivos.
func Init(level string) *zap.Logger {
	// Definir o nível de log.
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// Garantir que a pasta "logs" exista.
	logDir := "logs"
	logFile := filepath.Join(logDir, "app.log")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic("Falha ao criar a pasta de logs: " + err.Error())
	}

	// Configuração do Zap.
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapLevel),
		Encoding:    "json", // Pode usar "console" para texto mais legível.
		OutputPaths: []string{"stdout", logFile}, // Registrar no terminal e no arquivo.
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			MessageKey:     "msg",
			CallerKey:      "caller",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // Níveis em letras minúsculas.
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // Timestamp ISO8601.
			EncodeCaller:   zapcore.ShortCallerEncoder,    // Arquivo e linha curtos.
		},
	}

	// Criar o logger.
	logger, err := config.Build()
	if err != nil {
		panic("Falha ao criar o logger: " + err.Error())
	}

	return logger
}
