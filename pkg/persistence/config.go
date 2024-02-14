package persistence

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const connStringf string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

// DatabaseConfig is the database config structure
type DatabaseConfig struct {
	User               string        `mapstructure:"user"`
	Password           string        `mapstructure:"password"`
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	Name               string        `mapstructure:"name"`
	SSLMode            string        `mapstructure:"sslMode"`
	MaxOpenConnections int           `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int           `mapstructure:"maxIdleConnections"`
	ConnMaxLifetime    time.Duration `mapstructure:"connMaxLifetime"`
}

// GetConnString build a connection string for the database from DatabaseConfig
func (cfg DatabaseConfig) GetConnString() string {
	return fmt.Sprintf(connStringf, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
}

// LoadConfigFromYAML loads database configuration from a YAML file
func LoadConfigFromYAML(filepath string) (*DatabaseConfig, error) {
	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config DatabaseConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &config, nil
}
