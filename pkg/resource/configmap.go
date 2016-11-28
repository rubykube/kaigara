package resource

import (
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type ConfigMap map[interface{}]interface{}

func ParseConfigMap(configMapData string) ConfigMap {
	cm := make(ConfigMap)

	err := yaml.Unmarshal([]byte(configMapData), &cm)
	if err != nil {
		term.Error(err.Error())
	}

	data := cm["data"].(ConfigMap)
	values := data["values.yaml"].(ConfigMap)

	return values
}

func ApplyConfigMap(cm ConfigMap) {
	for key, value := range cm {
		viper.Set(key.(string), value)
	}
}
