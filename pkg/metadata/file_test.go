package metadata

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	file, _ := os.Create("metadata.yml")
	filename, _ := filepath.Abs(file.Name())

	file.WriteString("test: data")
	file.Close()

	Parse()

	os.Remove("metadata.yml")

	if viper.ConfigFileUsed() != filename {
		t.Fail()
	}

	if viper.Get("test") != "data" {
		t.Fail()
	}
}
