package log

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var (
	coreZapOptions       = []zap.Option{zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)}
	coreZapSlogOptions   = []zapslog.HandlerOption{zapslog.WithCaller(true), zapslog.AddStacktraceAt(slog.LevelError)}
	coreZapEncoderConfig = zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
)

func NewDevelopment() (*slog.Logger, *zap.Logger) {
	config := coreZapEncoderConfig
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeDuration = zapcore.StringDurationEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.Lock(os.Stderr), zap.DebugLevel)

	handler := zapslog.NewHandler(core, coreZapSlogOptions...)
	sLogger := slog.New(handler)

	zapOptions := append(coreZapOptions, zap.Development())
	zLogger := zap.New(core, zapOptions...)

	return sLogger, zLogger
}

func NewProduction() (*slog.Logger, *zap.Logger) {
	config := coreZapEncoderConfig
	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncodeTime = zapcore.EpochTimeEncoder
	config.EncodeDuration = zapcore.SecondsDurationEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.Lock(os.Stderr), zap.InfoLevel)

	handler := zapslog.NewHandler(core, coreZapSlogOptions...)
	sLogger := slog.New(handler)

	zLogger := zap.New(core, coreZapOptions...)

	return sLogger, zLogger
}

func Err(err error) slog.Attr {
	return slog.String("err", err.Error())
}
