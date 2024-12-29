package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

const EnvFilename = ".env"

func LoadEnv() (env *EnvironmentVariable, err error) {
	slog.Info("Load Env here")

	envFile := EnvFilename

	v := viper.New()

	v.SetConfigFile(envFile)
	err = v.ReadInConfig()
	if err != nil {
		slog.Error("Failed to load Env Config", "filename", envFile, "error", err)
	}

	v.AutomaticEnv()
	err = v.Unmarshal(&env)
	if err != nil {
		slog.Error("Failed to parse Env Config", "filename", envFile, "error", err)
	}

	return env, nil
}

type EnvironmentVariable struct {
	Database struct {
		PG_HOST     string `mapstructure:"PG_HOST"`
		PG_PORT     int    `mapstructure:"PG_PORT"`
		PG_NAME     string `mapstructure:"PG_NAME"`
		PG_USERNAME string `mapstructure:"PG_USERNAME"`
		PG_PASSWORD string `mapstructure:"PG_PASSWORD"`
		PG_TIMEZONE string `mapstructure:"PG_TIMEZONE"`
	} `mapstructure:"DATABASE"`
}
