package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logFile = "app.log"
const timeFormat string = "2006-01-02T15:04:05.999Z"

var zapLogger, _ = newLogger(zap.InfoLevel)

func Logger() *zap.Logger {
	return zapLogger
}

func newLogger(level zapcore.Level) (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "sourceLocation",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(timeFormat),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", logFile},
		ErrorOutputPaths: []string{"stderr", logFile},
	}
	return config.Build(zap.AddStacktrace(zap.FatalLevel))
}

// example usage
//		logger.Logger().Debug("this is a debug message") 	// Information useful to developers during debugging.
//		logger.Logger().Info("this is an info message") 	// Confirms that the application is working the way it is supposed to.
//		logger.Logger().Warn("this is a warn message") 	// Indicates a problem that can disturb the application in the future.
//		logger.Logger().Error("this is an error message") // An issue causing malfunctioning of one or more features.
//		logger.Logger().Fatal("this is a fatal message") 	// a serious issue that prevents the program from working.
//		logger.Logger().Panic("this is a panic message") // a critical issue that forces the program to stop.
