package api

import "github.com/arvph/ServerAndDB/storage"

// General instance for API server of REST application
type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	// Store config
	Storage *storage.Config
}

// функция для создания структуры Config и заполнение ее базовыми значениями
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(), // ссылается на изначально пустую структуру storage.Config
	}
}
