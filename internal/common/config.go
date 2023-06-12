package common

var (
	cfg Config
)

func SetConfig(c Config) {
	cfg = c
}

func GetConfig() Config {
	return cfg
}
