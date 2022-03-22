package logger

import (
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger
var Logger *zap.Logger

// SugarLogger
var SugarLogger *zap.SugaredLogger

// logConfig
type logConfig struct {
	LogPath    string
	LogLevel   string
	Compress   bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

// Init conf
func Init(opts ...Option) {
	lcf := defaultOption()
	for _, opt := range opts {
		opt(lcf)
	}
	initLogger(lcf)
}

// InitLogger Logger/SugarLogger
func initLogger(cfg *logConfig) {
	writeSyncer := zapcore.NewMultiWriteSyncer(getLogWriter(cfg), zapcore.AddSync(os.Stderr))
	encoder := getEncoder()
	var enab = new(zapcore.Level)
	err := enab.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		log.Fatal(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, enab)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(3))
	SugarLogger = Logger.Sugar()
}

func getLogWriter(lc *logConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   lc.LogPath + "server.log",
		MaxSize:    lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge:     lc.MaxAge,
		Compress:   lc.Compress,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func defaultOption() *logConfig {
	return &logConfig{
		LogPath:    "logs/",
		MaxSize:    20,
		Compress:   true,
		MaxAge:     100,
		MaxBackups: 1000,
		LogLevel:   "debug",
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "severity"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig) // Console Encoder-->NewJSONEncoder
}

func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	SugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	SugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	SugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	SugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	SugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	SugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	SugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	SugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	SugarLogger.Fatalf(template, args...)
}

func LoggerSync() error {
	return Logger.Sync()
}
