package metadata

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	// test without argument
	file, _ := os.Create("defaults.yml")
	filename, _ := filepath.Abs(file.Name())

	file.WriteString("test: data1")
	file.Close()

	Parse("")

	os.Remove("defaults.yml")

	if viper.ConfigFileUsed() != filename {
		t.Fail()
	}

	if viper.Get("test") != "data1" {
		t.Fail()
	}

	// test with argument
	testDir := "test/test/test/test" // more test :D
	testFile := testDir + "/test.yml"

	os.MkdirAll(testDir, 0755)
	file, _ = os.Create(testFile)
	filename, _ = filepath.Abs(file.Name())

	file.WriteString("test: data2")
	file.Close()

	Parse(testFile)

	os.Remove(testFile)
	os.Remove(testDir)

	if viper.ConfigFileUsed() != filename {
		t.Fail()
	}

	if viper.Get("test") != "data2" {
		t.Fail()
	}
}
