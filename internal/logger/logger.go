// Package logger – отвечает за логгирование
package logger

import (
	"go.uber.org/zap"
)

// Log – глобальная переменная, которая даёт доступ к логгеру
var Log *zap.SugaredLogger

// Initialize – функция для инициализации конфига
// Парсит уровень логирования, создаёт SugaredLogger и сеттит его в переменную Log
func Initialize(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	Log = zl.Sugar()
	return nil
}
