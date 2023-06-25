package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//	func CreateLogger() (*zap.Logger, error) {
//		logger, err := zap.NewProduction()
//		if err != nil {
//			return nil, err
//		}
//		return logger, nil
//	}
func CreateLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.LevelKey = "severity"
	return config.Build()
}
