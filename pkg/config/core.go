package config

import (
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/viper"
	"strings"
)

var cfg *Config

type Config struct {
	path []string
}

func setDefaults() {
	viper.SetDefault("home", "/opt/provision")
	viper.SetDefault("path", "/opt/provision/operations")
	viper.SetDefault("tmpl", "/opt/provision/resources")
	viper.SetDefault("env", "development")
}

func parseConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kaigara")
	viper.AddConfigPath("$HOME/.kaigara")

	err := viper.ReadInConfig()
	if err != nil {
		term.Say("No configuration found")
	}
}

func parsePath() []string {
	mp := ".:" + Get("path")
	return strings.Split(mp, ":")
}

func Init() {
	cfg = new(Config)
	setDefaults()
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
