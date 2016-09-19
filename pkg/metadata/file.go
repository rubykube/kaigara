package metadata

import (
	"github.com/mod/kaigara/pkg/config"
	"github.com/mod/kaigara/pkg/file"
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/viper"
)

var metaFile string

func SetFile() {
	if metaFile != "" {
		viper.SetConfigFile(metaFile)
	}
}

func metaPath() []string {
	metapath := []string{".", config.Get("home")}
	return metapath
}

func Parse() {
	viper.SetConfigName("metadata")
	for _, dir := range metaPath() {
		if file.Exists(dir + "/metadata.yml") {
			viper.AddConfigPath(dir)
			break
		}
	}
	if err := viper.ReadInConfig(); err == nil {
		term.Warning("Using metafile: " + viper.ConfigFileUsed())
	} else {
		term.Warning("No Metafile found")
	}
}
