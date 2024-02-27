package storage

type Config struct {
	DatabaseURI string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
