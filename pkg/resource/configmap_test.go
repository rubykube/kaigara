package resource

import (
	"testing"

	"github.com/spf13/viper"
)

var data = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: tmp
data:
  values.yaml:
    test1: 123
    test2: kaigara
`

func TestParseConfigMap(t *testing.T) {
	cm := ParseConfigMap(data)
	if cm["test1"] != 123 || cm["test2"] != "kaigara" {
		t.Fail()
	}
}

func TestApplyConfigMap(t *testing.T) {
	cm := ParseConfigMap(data)
	ApplyConfigMap(cm)

	eq := false

	for configMapKey, configMapValue := range cm {
		for viperKey, viperValue := range viper.AllSettings() {
			if viperKey == configMapKey && viperValue == configMapValue {
				eq = true
			}
		}
	}

	if !eq {
		t.Fail()
	}
}
