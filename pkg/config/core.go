package config

import (
	"strings"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	path []string
}

func parseConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kaigara")

	viper.ReadInConfig()
}

func parsePath() []string {
	mp := ".:" + Get("path")
	return strings.Split(mp, ":")
}

func Init() {
	cfg = new(Config)
	parseConfig()
	viper.SetEnvPrefix("kaigara")
	viper.AutomaticEnv()
}

func Get(name string) string {
	return viper.GetString(name)
}

func GetPath() []string {
	if cfg.path == nil {
		cfg.path = parsePath()
	}
	return cfg.path
}
