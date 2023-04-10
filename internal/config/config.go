package config

type Config struct {
	DBPath string
	Port   string
}

func NewConfig() *Config {
	return &Config{
		DBPath: "forum.db",
		Port:   "8080",
	}
}
