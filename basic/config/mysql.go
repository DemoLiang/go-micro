package config

import "time"

type MysqlConfig interface {
	GetUrl() string
	GetEnabled() bool
	GetMalIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() time.Duration
}

type defaultMysqlConfig struct {
	URL               string        `json:"url"`
	Enabled            bool          `json:"enabled"`
	MaxIdleConnection int           `json:"max_idle_connection"`
	MaxOpenConnection int           `json:"max_open_connection"`
	ConnMaxLifetime   time.Duration `json:"conn_max_lifetime"`
}

func (m defaultMysqlConfig) GetUrl() string {
	return m.URL
}

func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enabled
}

func (m defaultMysqlConfig) GetMalIdleConnection() int {
	return m.MaxIdleConnection
}

func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

func (m defaultMysqlConfig) GetConnMaxLifetime() time.Duration {
	return m.ConnMaxLifetime
}
