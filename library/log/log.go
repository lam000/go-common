package log

import (
	"context"
	"fmt"
	"github.com/lam000/go-common/library/trace"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	FileName   string
	RotateSize int
	MaxBackups int
	MaxAge     int
	LocalTime  bool
	Compress   bool
	Level      string
}

var logger *zap.Logger

func Init(conf *Config) {
	writerSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.FileName,
		MaxSize:    conf.RotateSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
	})

	var atomicLevel zap.AtomicLevel
	if err := atomicLevel.UnmarshalText([]byte(conf.Level)); err != nil {
		panic(fmt.Errorf("invalid logger level %s", conf.Level))
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writerSyncer, atomicLevel)
	logger = zap.New(core, zap.AddCaller())
}

func RawInstance() *zap.Logger {
	return logger
}

func Instance() *zap.SugaredLogger {
	return RawInstance().Sugar()
}

func CtxRawInstance(ctx context.Context) *zap.Logger {
	traceID, ok := ctx.Value(trace.TID).(string)
	if !ok {
		return logger
	}

	return logger.With(zap.String(trace.TID, traceID))
}

func CtxInstance(ctx context.Context) *zap.SugaredLogger {
	return CtxRawInstance(ctx).Sugar()
}
