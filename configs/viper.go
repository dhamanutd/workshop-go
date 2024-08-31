package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION"`

	PostgreSQLConfig PostgreSQLConfig `mapstructure:"POSTGRESQL_DATABASE"`
}

type Secret struct {
	PostgreSQLSecret PostgreSQLSecret `mapstructure:"POSTGRESQL_DATABASE"`
}

type PostgreSQLConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type PostgreSQLSecret struct {
	Database string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

var env string

func LoadConfig(configPath, secretPath string) (*Config, *Secret, error) {
	// Load Config
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	viper.Reset()

	// Load Secret
	viper.SetConfigFile(secretPath)
	var secret Secret
	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("error reading secret file: %w", err)
	}
	if err := viper.Unmarshal(&secret); err != nil {
		return nil, nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	env = config.Environment

	return &config, &secret, nil
}
