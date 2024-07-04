package logr

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync/atomic"
)

var logger atomic.Value

type Logger struct {
	logger *zap.Logger
}

func Default() *Logger {
	l, ok := logger.Load().(*Logger)
	if !ok {
		l = New()
		logger.Store(l)
	}
	return l
}

func Set(l *Logger) {
	logger.Store(l)
}

func New(opts ...Option) *Logger {

	c := config{
		Level:       zap.NewAtomicLevel(),
		WriteSyncer: zapcore.AddSync(os.Stdout),
	}

	for _, opt := range opts {
		opt(&c)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if c.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(encoder, c.WriteSyncer, c.Level)

	return &Logger{
		logger: zap.New(core,
			zap.AddCallerSkip(1),
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
			zap.Development(),
			zap.ErrorOutput(c.WriteSyncer),
		).WithOptions(c.Options...),
	}
}

func (l *Logger) Sync() {
	_ = l.logger.Sync()
}

func Sync() {
	Default().Sync()
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return l
	}
	clone := *l
	clone.logger = clone.logger.With(fields...)
	return &clone
}

func With(fields ...zap.Field) *Logger {
	return Default().With(fields...)
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	return l.With(FieldFromContext(ctx)...)
}

func WithContext(ctx context.Context) *Logger {
	return Default().WithContext(ctx)
}

func (l *Logger) WithOption(opts ...zap.Option) *Logger {
	if len(opts) == 0 {
		return l
	}
	clone := *l
	clone.logger = clone.logger.WithOptions(opts...)
	return &clone
}

func WithOption(opts ...zap.Option) *Logger {
	return Default().WithOption(opts...)
}

func FieldFromContext(ctx context.Context) []zap.Field {
	spanCtx := trace.SpanContextFromContext(ctx)
	var fields []zap.Field
	if spanCtx.HasTraceID() {
		fields = append(fields, zap.String("trace", spanCtx.TraceID().String()))
	}
	if spanCtx.HasSpanID() {
		fields = append(fields, zap.String("span", spanCtx.SpanID().String()))
	}
	return fields
}
