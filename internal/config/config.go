package config

import (
	"github.com/spf13/viper"

	"github.com/art-injener/otus/internal/logger"
)

const DebugLevel = "debug"

type Config struct {
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	ServerPort int    `mapstructure:"PORT"`
	UseMocks   bool   `mapstructure:"USE_MOCK"`
	DBConfig
	Log *logger.Logger
}

func (c *Config) IsDebug() bool {
	return c.LogLevel == DebugLevel
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     uint16 `mapstructure:"DB_PORT"`
	NameDB   string `mapstructure:"DB_DATABASE"`
	User     string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("DB_HOST", &cfg.DBConfig.Host); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("DB_PORT", &cfg.DBConfig.Port); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("DB_USERNAME", &cfg.DBConfig.User); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("DB_PASSWORD", &cfg.DBConfig.Password); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("DB_DATABASE", &cfg.DBConfig.NameDB); err != nil {
		return nil, err
	}

	return &cfg, nil
}
