package hklog

import (
	"go.hikit.io/hkctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_Log *zap.SugaredLogger

	_Ctx       = hkctx.Background()
	_TraceKey  = "trace_id"
	_LoggerKey = "logger_name"
)

func init() {
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	_Ctx = hkctx.WithValue(hkctx.Background(), _TraceKey, "no id")
	_Ctx = hkctx.WithValue(_Ctx, _LoggerKey, "hklog")
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	c := zap.Config{
		Level:             atom,
		Development:       true,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		DisableStacktrace: true,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
	}
	if l, err := c.Build(zap.AddCallerSkip(1)); err != nil {
		panic(err)
	} else {
		_Log = l.Sugar()
	}
}

func ReplaceCtx(ctx hkctx.Ctx) {
	_Ctx = ctx
}

func ReplaceLog(log *zap.SugaredLogger) {
	_Log = log
}

func ReplaceTraceKey(name string) {
	_TraceKey = name
}

func ReplaceLoggerKey(name string) {
	_LoggerKey = name
}

func TraceKey() string {
	return _TraceKey
}

func LoggerKey() string {
	return _LoggerKey
}

func Debug(ctx hkctx.Ctx, msg string, keyAndValues ...interface{}) {
	keyAndValues = append(keyAndValues,
		zap.Any(_TraceKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_TraceKey)),
		zap.Any(_LoggerKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_LoggerKey)),
	)
	_Log.Debugw(msg, keyAndValues...)
}

func Error(ctx hkctx.Ctx, msg string, keyAndValues ...interface{}) {
	keyAndValues = append(keyAndValues,
		zap.Any(_TraceKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_TraceKey)),
		zap.Any(_LoggerKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_LoggerKey)),
	)
	_Log.Errorw(msg, keyAndValues...)
}

func Info(ctx hkctx.Ctx, msg string, keyAndValues ...interface{}) {
	keyAndValues = append(keyAndValues,
		zap.Any(_TraceKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_TraceKey)),
		zap.Any(_LoggerKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_LoggerKey)),
	)
	_Log.Infow(msg, keyAndValues...)
}

func Warn(ctx hkctx.Ctx, msg string, keyAndValues ...interface{}) {
	keyAndValues = append(keyAndValues,
		zap.Any(_TraceKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_TraceKey)),
		zap.Any(_LoggerKey, hkctx.DefaultCtx(ctx, Ctx()).Value(_LoggerKey)),
	)
	_Log.Warnw(msg, keyAndValues...)
}

func Ctx() hkctx.Ctx {
	return _Ctx
}
