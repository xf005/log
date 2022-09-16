package log

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l        *Logger
	outWrite zapcore.WriteSyncer // IO输出
	once     sync.Once
)

type Logger struct {
	*zap.Logger
	opts      *Options
	zapConfig zap.Config
}

func NewLogger(opts ...LogOptions) {
	once.Do(func() {
		l = &Logger{
			opts: newOptions(opts...),
		}
		l.loadCfg()
		l.initLog()
		l.Info("zap plugin initializing completed")
	})
}

func WithContext(ctx context.Context) *zap.Logger {
	if l == nil {
		fmt.Println("zap plugin initializing...")
		NewLogger()
	}
	log, ok := ctx.Value(l.opts.CtxKey).(*zap.Logger)
	if ok {
		return log
	}
	return l.Logger
}

func AddCtx(ctx context.Context, field ...zap.Field) (context.Context, *zap.Logger) {
	if l == nil {
		fmt.Println("zap plugin initializing...")
		NewLogger()
	}
	log := l.With(field...)
	ctx = context.WithValue(ctx, l.opts.CtxKey, log)
	return ctx, log
}

func (l *Logger) initLog() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}
func (l *Logger) GetLevel() (level zapcore.Level) {
	switch strings.ToLower(l.opts.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (l *Logger) loadCfg() {
	l.zapConfig = zap.NewProductionConfig()
	l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
}

func (l *Logger) setSyncers() {
	outWrite = zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", l.opts.LogFileDir, l.opts.Name),
		MaxSize:    l.opts.MaxSize,
		MaxAge:     l.opts.MaxAge,
		MaxBackups: 10,
		Compress:   true,
		LocalTime:  true,
	})
	return
}

func (l *Logger) cores() zap.Option {
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l.GetLevel()
	})
	var cores []zapcore.Core
	encoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	if l.opts.Dev {
		cores = append(cores,
			zapcore.NewCore(encoder, outWrite, priority),                // 写入文件
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), priority), // 写入控制台
		)
	} else {
		cores = append(cores,
			zapcore.NewCore(encoder, outWrite, priority),
		)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

// timeEncoder 日志输出时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
