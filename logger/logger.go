package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	// Creating a custom config so we can see what we wish in the log
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	// //initiating vanila logger
	// log, err = zap.NewProduction()

	// // If we want to see the functions which are calling the logger, we can do the following
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))

	// // logger with custom config
	// log, err = config.Build()

	// // we can display function that is calling the logger similarly as before
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info ... Using this, we avoid displaying the actual function which is calling the logger
// so all the logs will now shoow logger/logger.
// The other method was to call log.Info in the main function
// But that would have shown the function which is calling the logger.
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Other helper functions like Info

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
