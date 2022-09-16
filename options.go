package log

import "path/filepath"

type Options struct {
	Dev     bool   // 是否开发模式
	Level   string // 日志打印等级
	CtxKey  string // 通过 ctx 传递 log 对象
	Name    string // APP名字
	MaxSize int    // 文件最大多少M进行分割
	MaxAge  int    // 文件保留天数
	FileDir string // 日志路径
}

type Option func(*Options)

func newOptions(opts ...Option) *Options {
	opt := &Options{
		Dev:     true,
		Level:   "debug",
		CtxKey:  "log-ctx",
		Name:    "test",
		MaxSize: 500,
		MaxAge:  30,
	}
	opt.FileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
	opt.FileDir += "./logs/"
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func SetDev(dev bool) Option {
	return func(options *Options) {
		options.Dev = dev
	}
}

func SetFileDir(fileDir string) Option {
	return func(options *Options) {
		options.FileDir = fileDir
	}
}

func SetName(name string) Option {
	return func(options *Options) {
		options.Name = name
	}
}

func SetMaxSize(maxSize int) Option {
	return func(options *Options) {
		options.MaxSize = maxSize
	}
}

func SetMaxAge(maxAge int) Option {
	return func(options *Options) {
		options.MaxAge = maxAge
	}
}

func SetLevel(level string) Option {
	return func(options *Options) {
		options.Level = level
	}
}

func SetCtxKey(ctxKey string) Option {
	return func(options *Options) {
		options.CtxKey = ctxKey
	}
}
