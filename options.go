package log

import "path/filepath"

type Options struct {
	Dev        bool   // 是否开发模式
	Level      string // 日志打印等级
	CtxKey     string // 通过 ctx 传递 log 对象
	Name       string // APP名字
	MaxSize    int    // 文件最大多少M进行分割
	MaxAge     int    // 文件保留天数
	LogFileDir string // 日志路径
}

type LogOptions func(*Options)

func newOptions(opts ...LogOptions) *Options {
	opt := &Options{
		Dev:     true,
		Level:   "debug",
		CtxKey:  "log-ctx",
		Name:    "test",
		MaxSize: 500,
		MaxAge:  30,
	}
	opt.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
	opt.LogFileDir += "./logs/"
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func SetDev(dev bool) LogOptions {
	return func(options *Options) {
		options.Dev = dev
	}
}

func SetLogFileDir(logFileDir string) LogOptions {
	return func(options *Options) {
		options.LogFileDir = logFileDir
	}
}

func SetName(name string) LogOptions {
	return func(options *Options) {
		options.Name = name
	}
}

func SetMaxSize(maxSize int) LogOptions {
	return func(options *Options) {
		options.MaxSize = maxSize
	}
}

func SetMaxAge(maxAge int) LogOptions {
	return func(options *Options) {
		options.MaxAge = maxAge
	}
}

func SetLevel(level string) LogOptions {
	return func(options *Options) {
		options.Level = level
	}
}

func SetCtxKey(ctxKey string) LogOptions {
	return func(options *Options) {
		options.CtxKey = ctxKey
	}
}
