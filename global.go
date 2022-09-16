package log

import (
	"context"
	"go.uber.org/zap"
)

// Debugs 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 debug 级别的日志
// logging.Debugs(nil, "abc", 123)
func Debugs(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Debug(args...)
}

// Infos 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 info 级别的日志
func Infos(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Info(args...)
}

// Warns 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 warn 级别的日志
func Warns(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Warn(args...)
}

// Errors 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 Error 级别的日志
func Errors(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Error(args...)
}

// Panics 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 Panic 级别的日志
func Panics(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Panic(args...)
}

// Fatals 尝试从 Context 中获取带 trace id 的 sugared logger 来记录 Fatal 级别的日志
func Fatals(c context.Context, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Fatal(args...)
}

// Debugf 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 debug 级别的日志
// logging.Debugf(nil, "str:%s", "abd")
func Debugf(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Debugf(template, args...)
}

// Infof 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 info 级别的日志
func Infof(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Infof(template, args...)
}

// Warnf 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 warn 级别的日志
func Warnf(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Warnf(template, args...)
}

// Errorf 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 error 级别的日志
func Errorf(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Errorf(template, args...)
}

// Panicf 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 panic 级别的日志
func Panicf(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Panicf(template, args...)
}

// Fatalf 尝试从 Context 中获取带 trace id 的 sugared logger 来模板字符串记录 fatal 级别的日志
func Fatalf(c context.Context, template string, args ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Fatalf(template, args...)
}

// Debugw 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 debug 级别的日志
// logging.Debugw(nil, "msg", "k1", "v1", "k2", "v2")
func Debugw(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Debugw(msg, keysAndValues...)
}

// Infow 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 info 级别的日志
func Infow(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Infow(msg, keysAndValues...)
}

// Warnw 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 warn 级别的日志
func Warnw(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Warnw(msg, keysAndValues...)
}

// Errorw 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 error 级别的日志
func Errorw(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Errorw(msg, keysAndValues...)
}

// Panicw 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 panic 级别的日志
func Panicw(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Panicw(msg, keysAndValues...)
}

// Fatalw 尝试从 Context 中获取带 trace id 的 sugared logger 来 kv 记录 fatal 级别的日志
func Fatalw(c context.Context, msg string, keysAndValues ...interface{}) {
	slogger := WithContext(c).WithOptions(zap.AddCallerSkip(1)).Sugar()
	defer slogger.Sync()
	slogger.Fatalw(msg, keysAndValues...)
}

// Debug 尝试从 Context 中获取带 trace id 的 logger 记录 debug 级别的日志
func Debug(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Debug(msg, fields...)
}

// Info 尝试从 Context 中获取带 trace id 的 logger 记录 info 级别的日志
func Info(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Info(msg, fields...)
}

// Warn 尝试从 Context 中获取带 trace id 的 logger 记录 warn 级别的日志
func Warn(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Warn(msg, fields...)
}

// Error 尝试从 Context 中获取带 trace id 的 logger 记录 error 级别的日志
func Error(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Error(msg, fields...)
}

// Panic 尝试从 Context 中获取带 trace id 的 logger 记录 panic 级别的日志
func Panic(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Panic(msg, fields...)
}

// Fatal 尝试从 Context 中获取带 trace id 的 logger 记录 fatal 级别的日志
func Fatal(c context.Context, msg string, fields ...zap.Field) {
	logger := WithContext(c).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
	logger.Fatal(msg, fields...)
}
