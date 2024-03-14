package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const logFile = "app.log"

var Sugar *zap.SugaredLogger

func InitLogger() {

	// Setting the logging level.
	logLevel := zap.InfoLevel

	// Encoder configuration for both console and file.
	// This configuration specifies how logs should be formatted and encoded.
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // Encode the level in Capital Case.
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // Encode time in ISO8601 format.
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // Include file name and line number.
	}

	// Creating a console encoder with the above configuration for human-readable output.
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Creating a file encoder; also using the console encoder for text format.
	// You can use a JSON encoder here if you prefer JSON format in files.
	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Core for console output, filtering logs according to the log level.
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	// Core for file logging, using the file encoder.
	// Opens or creates the file for appending logs.
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // Handling errors opening or creating the log file.
	}
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), logLevel)

	// Combining both cores to log both to console and file with the same log level.
	core := zapcore.NewTee(consoleCore, fileCore)

	// Creating the logger with the combined core and adding caller information (file and line number).
	logger := zap.New(core, zap.AddCaller())

	// Assigning the sugared logger to the global variable.
	// The sugared logger offers a more flexible API for logging.
	Sugar = logger.Sugar()
}

// example usage
//		sugar.Debug("this is a debug message") 	// Information useful to developers during debugging.
//		sugar.Info("this is an info message") 	// Confirms that the application is working the way it is supposed to.
//		sugar.Warn("this is a warn message") 	// Indicates a problem that can disturb the application in the future.
//		sugar.Error("this is an error message") // An issue causing malfunctioning of one or more features.
//		sugar.Fatal("this is a fatal message") 	// a serious issue that prevents the program from working.
