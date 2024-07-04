package logr

import (
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()
	logs, cleanup, err := initLogger("bottle-wallet")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	Infoc(ctx, "debug", zap.String("times", time.Now().String()))
	slog := With(zap.String("space", "slog"))
	slog.Infoc(ctx, "debug", zap.String("times", time.Now().String()))
	Error("error", zap.Error(fmt.Errorf("test error")))
	logs.Infoc(ctx, "debug", zap.String("times", time.Now().String()))
}

func initLogger(appName string) (*Logger, func(), error) {
	logLevel, ok := os.LookupEnv("KUBE_LOG_LEVEL")
	if !ok {
		logLevel = "debug"
	}
	logPath, ok := os.LookupEnv("KUBE_LOG_PATH")
	if !ok {
		logPath = "log"
	}
	fileName := fmt.Sprintf("%s/%s/log", logPath, appName)
	writer, err := rotatelogs.New(
		fileName+".%Y%m%d",                        // 日志文件名
		rotatelogs.WithLinkName(fileName),         // 创建软连接
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 日志保存最大时长
	)
	if err != nil {
		return nil, func() {}, err
	}
	logger := New(
		SetLevel(logLevel),
		SetFormat("json"),
		SetWriter(zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer))),
		SetOptions(zap.Fields(zap.String("service", appName))),
	)
	Set(logger)
	return logger, func() {
		logger.Sync()
		_ = writer.Close()
	}, nil
}
