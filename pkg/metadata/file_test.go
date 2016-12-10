package metadata

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, _ := os.Create("defaults.yml")

	file.WriteString("test: data")
	file.Close()

	Parse()

	os.Remove("defaults.yml")
}
