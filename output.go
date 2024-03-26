package logr

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Debug(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Infoc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Info(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Warn(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Error(msg, append(fields, FieldFromContext(ctx)...)...)
}

func Debug(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugf(format, args...)
}

func Debugc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugc(ctx, msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infof(format, args...)
}

func Infoc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Infoc(ctx, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Warnf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnf(format, args...)
}

func Warnc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnc(ctx, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Errorf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorf(format, args...)
}

func Errorc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorc(ctx, msg, fields...)
}
