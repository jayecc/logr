package logr

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func w(args []interface{}) (fields []zap.Field) {
	for i := 0; i < len(args)/2; i++ {
		fields = append(fields, zap.Any(fmt.Sprint(args[i*2]), args[i*2+1]))
	}
	return
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugw(msg string, args ...interface{}) {
	l.logger.Debug(msg, w(args)...)
}

func (l *Logger) Debugc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Debug(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Debugcf(ctx context.Context, format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...), FieldFromContext(ctx)...)
}

func (l *Logger) Debugcw(ctx context.Context, msg string, args ...interface{}) {
	l.logger.Debug(msg, append(w(args), FieldFromContext(ctx)...)...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Infow(msg string, args ...interface{}) {
	l.logger.Info(msg, w(args)...)
}

func (l *Logger) Infoc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Info(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Infocf(ctx context.Context, format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...), FieldFromContext(ctx)...)
}

func (l *Logger) Infocw(ctx context.Context, msg string, args ...interface{}) {
	l.logger.Info(msg, append(w(args), FieldFromContext(ctx)...)...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnw(msg string, args ...interface{}) {
	l.logger.Warn(msg, w(args)...)
}

func (l *Logger) Warnc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Warn(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Warncf(ctx context.Context, format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...), FieldFromContext(ctx)...)
}

func (l *Logger) Warncw(ctx context.Context, msg string, args ...interface{}) {
	l.logger.Warn(msg, append(w(args), FieldFromContext(ctx)...)...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Errorw(msg string, args ...interface{}) {
	l.logger.Error(msg, w(args)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorc(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Error(msg, append(fields, FieldFromContext(ctx)...)...)
}

func (l *Logger) Errorcf(ctx context.Context, format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...), FieldFromContext(ctx)...)
}

func (l *Logger) Errorcw(ctx context.Context, msg string, args ...interface{}) {
	l.logger.Error(msg, append(w(args), FieldFromContext(ctx)...)...)
}

func Debug(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugf(format, args...)
}

func Debugw(msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugw(msg, args...)
}

func Debugc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugc(ctx, msg, fields...)
}

func Debugcf(ctx context.Context, format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugcf(ctx, format, args...)
}

func Debugcw(ctx context.Context, msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugcw(ctx, msg, args...)
}

func Info(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infof(format, args...)
}

func Infow(msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infow(msg, args...)
}

func Infoc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Infoc(ctx, msg, fields...)
}

func Infocf(ctx context.Context, format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infocf(ctx, format, args...)
}

func Infocw(ctx context.Context, msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infocw(ctx, msg, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Warnf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnf(format, args...)
}

func Warnw(msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnw(msg, args...)
}

func Warnc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnc(ctx, msg, fields...)
}

func Warncf(ctx context.Context, format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warncf(ctx, format, args...)
}

func Warncw(ctx context.Context, msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warncw(ctx, msg, args...)
}

func Error(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Errorf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorf(format, args...)
}

func Errorw(msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorw(msg, args...)
}

func Errorc(ctx context.Context, msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorc(ctx, msg, fields...)
}

func Errorcf(ctx context.Context, format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorcf(ctx, format, args...)
}

func Errorcw(ctx context.Context, msg string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorcw(ctx, msg, args...)
}
