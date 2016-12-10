package core

import (
	"os"

	"github.com/mod/kaigara/pkg/log"
)

type Config map[string]string
type EnvMap map[string]string

var cfg Config
var envMap = EnvMap{
	"core.path.log":        "KAIGARA_LOG",
	"core.path.metadata":   "KAIGARA_METADATA",
	"core.path.resources":  "KAIGARA_RESOURCES",
	"core.path.operations": "KAIGARA_OPERATIONS",
}

func Init() {
	log.Init(0, os.Stderr)
	log.Debug("Initializing kaigara core")

	cfg = MakeConfig()
	ImportEnv()
	log.Debug("Using metapath: " + cfg["core.path.metadata"])
}

func Get(key string) string {
	return cfg[key]
}

func Set(key string, value string) {
	cfg[key] = value
}

func MakeConfig() Config {
	config := Config{
		"core.path.log":        "/var/log/kaigara",
		"core.path.metadata":   "/etc/kaigara/metadata",
		"core.path.resources":  "/etc/kaigara/resources",
		"core.path.operations": "/opt/kaigara/operations",
	}
	return config
}

func ImportEnv() {
	for key, val := range envMap {
		UpdateConfig(key, val)
	}
}

func UpdateConfig(key string, env string) {
	value := os.Getenv(env)
	if value != "" {
		cfg[key] = value
		log.Debugf("update config %s=%s", key, value)
	}
}
