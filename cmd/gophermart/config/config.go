package config

import (
	"flag"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Config Settings

type Settings struct {
	Address netAddress
}

func SetConfig() error {

	Config = Settings{
		Address: netAddress{
			Host: "localhost", Port: 8080},
	}

	flag.Var(&Config.Address, "a", "Net address host:port")
	flag.Parse()

	if envAddr := os.Getenv("ADDRESS"); envAddr != "" {
		err := Config.Address.Set(envAddr)
		if err != nil {
			return err
		}
	}

	logger.LogNoSugar.Debug("Config", zap.Inline(Config))
	return nil
}

func (s Settings) MarshalLogObject(encoder zapcore.ObjectEncoder) error {

	err := encoder.AddObject("Address", &s.Address)
	if err != nil {
		return err
	}
	return err
}
