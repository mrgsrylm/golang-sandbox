package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ProductionEnvironment  = "production"
	StagingEnvironment     = "development"
	DevelopmentEnvironment = "staging"
)

// Default production is false if not call func
func New(environment string) {
	var conf zap.Config
	conf = ConfigProductionLogger()

	if environment != ProductionEnvironment {
		conf = zap.NewDevelopmentConfig()
	}

	conf.DisableStacktrace = true
	log, err := conf.Build()
	if err != nil {
		panic(err)
	}

	logger = log.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

/**
 ** NewProductionConfig is a reasonable production logging configuration.
 ** Logging is enabled at InfoLevel and above.
 **	It uses a CONSOLE encoder, writes to standard error, and enables sampling.
 ** Stacktraces are automatically included on logs of ErrorLevel and above.
**/
func ConfigProductionLogger() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    ConfigProductionEncoder(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// NewProductionEncoder returns an opinionated EncoderConfig for production environments.
func ConfigProductionEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
