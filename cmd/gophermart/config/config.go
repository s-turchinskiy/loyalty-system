package config

import (
	"flag"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Settings struct {
	Address       netAddress
	Database      database
	AccrualSystem string
}

func GetConfig() (*Settings, error) {

	config := &Settings{
		Address: netAddress{
			Host: "localhost", Port: 8080},
	}

	flag.Var(&config.Address, "a", "Net address host:port")
	flag.Var(&config.Database, "d", "path to database")
	flag.StringVar(&config.AccrualSystem, "r", "", "address of the accrual calculation system")
	flag.Parse()

	if envAddr := os.Getenv("RUN_ADDRESS"); envAddr != "" {
		err := config.Address.Set(envAddr)
		if err != nil {
			return nil, err
		}
	}

	if value := os.Getenv("DATABASE_URI"); value != "" {
		err := config.Database.Set(value)
		if err != nil {
			return nil, err
		}
	}

	if value := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); value != "" {
		config.AccrualSystem = value
	}

	logger.LogNoSugar.Info("Config", zap.Inline(config))
	return config, nil
}

func (s Settings) MarshalLogObject(encoder zapcore.ObjectEncoder) error {

	err := encoder.AddObject("Address", &s.Address)
	if err != nil {
		return err
	}
	return err
}
