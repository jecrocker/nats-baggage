package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(pack string) *zap.SugaredLogger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/output.log",
		MaxSize:    500,
		MaxBackups: 50,
		MaxAge:     1,
	})

	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(conf), w, zap.DebugLevel)
	stdout := zapcore.NewCore(zapcore.NewConsoleEncoder(conf), zapcore.AddSync(os.Stdout), zap.DebugLevel)
	logger := zap.New(zapcore.NewTee(core, stdout))

	sugar := logger.With(
		getHostname(),
		getApplication(),
		getPackage(pack),
	).Sugar()

	return sugar
}

func getHostname() zapcore.Field {
	hostname, err := os.Hostname()
	if err != nil {
		return zap.Skip()
	}
	return zap.String("hostname", hostname)
}

func getApplication() zapcore.Field {
	return zap.String("application", "baggage")
}

func getPackage(pack string) zapcore.Field {
	return zap.String("package", pack)
}
