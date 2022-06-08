package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapConfig struct {
	AppName     string        `mapstructure:"app_name"`
	Environment string        `mapstructure:"environment"` // local, development, staging, production
	Level       zapcore.Level `mapstructure:"level"`       // debug: -1, info: 0, ...
}

var Logger *zap.SugaredLogger

func InitZapLogger(config *ZapConfig) error {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	syncSlice := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
	if config.Environment == "production" {
		f, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		syncSlice = append(syncSlice, zapcore.AddSync(f))
	}

	syncer := zapcore.NewMultiWriteSyncer(syncSlice...)
	core := zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(config.Level))

	Logger = zap.New(core).With(zap.String("app", config.AppName)).Sugar()

	return nil
}
