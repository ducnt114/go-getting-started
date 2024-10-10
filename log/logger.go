package log

import (
	"context"
	"go-getting-started/enum"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func init() {
	ops := []zap.Option{
		zap.AddCallerSkip(2),
	}
	logger, _ = zap.NewProduction(ops...)
	zap.ReplaceGlobals(logger)
}

func getRequestID(ctx context.Context) string {
	id, ok := ctx.Value(enum.RequestIdKey).(string)
	if !ok {
		return ""
	}
	return id
}

func expand(ctx context.Context, data ...interface{}) []interface{} {
	if ctx == nil {
		return data
	}
	requestID := getRequestID(ctx)
	if requestID != "" {
		data = append(data, "request_id", requestID)
	}
	return data
}

func callLogFunc(logFn func(msg string, kv ...interface{}), msg string, data ...interface{}) {
	if len(data) == 0 {
		logFn(msg)
		return
	}
	logFn(msg, data...)
}

func Debugw(ctx context.Context, msg string, data ...interface{}) {
	callLogFunc(zap.S().Debugw, msg, expand(ctx, data...)...)
}

func Infow(ctx context.Context, msg string, data ...interface{}) {
	callLogFunc(zap.S().Infow, msg, expand(ctx, data...)...)
}

func Warnw(ctx context.Context, msg string, data ...interface{}) {
	callLogFunc(zap.S().Warnw, msg, expand(ctx, data...)...)
}

func Errorw(ctx context.Context, msg string, data ...interface{}) {
	callLogFunc(zap.S().Errorw, msg, expand(ctx, data...)...)
}
