package persistence

import (
	"fmt"
	"time"
)

const connStringf string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

// DatabaseConfig missing godoc
type DatabaseConfig struct {
	User               string
	Password           string
	Host               string
	Port               string
	Name               string
	SSLMode            string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    time.Duration
}

// GetConnString missing godoc
func (cfg DatabaseConfig) GetConnString() string {
	return fmt.Sprintf(connStringf, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
}
