package config

import (
	"os"
)

var cfg *Config

type Config struct {
	metapath string
	color    bool
}

func Init() {
	cfg = new(Config)
	cfg.metapath = readPath()
}

func GetPath() string {
	return cfg.metapath
}

func readPath() string {
	path := os.Getenv("KAIGARA_METADATA")
	if path != "" {
		return path
	}
	return "/etc/kaigara"
}
