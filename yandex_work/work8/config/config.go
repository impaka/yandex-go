package config

type Config struct {
	DefaultBalance float64
}

func NewConfig() *Config {
	return &Config{
		DefaultBalance: 1000,
	}
}
