package metadata

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mod/kaigara/pkg/core"
	"github.com/stretchr/testify/assert"
)

func MkTempDir(t *testing.T) string {
	cmd := exec.Command("mktemp", "-d")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(out.String())
}

func DumpFile(t *testing.T, path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	file.WriteString(content)
	file.Close()
}

func TestParse(t *testing.T) {
	tmpdir := MkTempDir(t)
	DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")

	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "they should be equal")

	os.RemoveAll(tmpdir)
}

func TestParseWithMerge(t *testing.T) {
	tmpdir := MkTempDir(t)
	DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")
	DumpFile(t, fmt.Sprintf("%s/additionals.yml", tmpdir), "info: data")

	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data", "info": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual, "they should be equal")

	os.RemoveAll(tmpdir)
}
