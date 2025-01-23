package log

import (
	"github.com/iter-x/iter-x/internal/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(isDev bool, cfg *conf.Log) (*zap.SugaredLogger, error) {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		level = zap.NewAtomicLevel()
	}
	zapCfg := zap.Config{
		Level:             level,
		Development:       isDev,
		DisableCaller:     cfg.DisableCaller,
		DisableStacktrace: cfg.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: cfg.Format,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "ns",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel: func() zapcore.LevelEncoder {
				if cfg.EnableColor {
					return zapcore.CapitalColorLevelEncoder
				}
				return zapcore.LowercaseLevelEncoder
			}(),
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{cfg.Output},
		ErrorOutputPaths: []string{cfg.Output},
	}
	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
