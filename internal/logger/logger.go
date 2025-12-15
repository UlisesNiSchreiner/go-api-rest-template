package logger

import "go.uber.org/zap"

type Logger struct{ *zap.Logger }

func New(env string) *Logger {
	cfg := zap.NewProductionConfig()
	if env == "dev" || env == "local" {
		cfg = zap.NewDevelopmentConfig()
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &Logger{l}
}

func String(key, val string) zap.Field { return zap.String(key, val) }
func Int(key string, val int) zap.Field { return zap.Int(key, val) }
func Err(err error) zap.Field { return zap.Error(err) }
