package metadata

import (
	"github.com/mod/kaigara/pkg/config"
	"github.com/mod/kaigara/pkg/file"
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/viper"
	"os"
	"strings"
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
	viper.SetConfigName("defaults")
	for _, dir := range metaPath() {
		if file.Exists(dir + "/defaults.yml") {
			viper.AddConfigPath(dir)
			break
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		term.Warning("No Metafile found")
	}
	parseEnv()
}

func parseEnv() {
	for _, e := range os.Environ() {
		v := strings.Split(e, "=")
		viper.Set(v[0], v[1])
	}
}

func ParseRemote() {
	var (
		provider      = config.Get("remote.provider")
		address       = config.Get("remote.address")
		path          = config.Get("remote.path")
		configType    = config.Get("remote.configType")
		secretKeyring = config.Get("remote.secretKeyring")
	)

	if secretKeyring == "" {
		viper.AddRemoteProvider(provider, address, path)
	} else {
		viper.AddSecureRemoteProvider(provider, address, path, secretKeyring)
	}

	viper.SetConfigType(configType)

	err := viper.ReadRemoteConfig()
	if err != nil {
		term.Error(err.Error())
	}
}
