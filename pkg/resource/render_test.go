package resource

import (
	"testing"

	"github.com/spf13/viper"
)

var testRenderData = `
test1: 123
test2: kaigara
`

func TestParseTemplate(t *testing.T) {
	viper.Set("test1", 123)
	viper.Set("test2", "kaigara")

	result := ParseTemplate("{{.test1}} {{.test2}}", viper.AllSettings())

	if result != "123 kaigara" {
		t.Fail()
	}
}
