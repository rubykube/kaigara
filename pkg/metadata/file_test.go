package metadata

import (
	"fmt"
	"os"
	"testing"

	"github.com/mod/kaigara/pkg/test"
	"github.com/mod/kaigara/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tmpdir := test.MkTempDir(t)
	test.DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")

	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	os.RemoveAll(tmpdir)
}

func TestParseWithMerge(t *testing.T) {
	tmpdir := test.MkTempDir(t)
	test.DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")
	test.DumpFile(t, fmt.Sprintf("%s/additionals.yaml", tmpdir), "info: data")

	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data", "info": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	os.RemoveAll(tmpdir)
}

func TestParseWithMergeNested(t *testing.T) {
	tmpdir := test.MkTempDir(t)
	os.Mkdir(tmpdir + "/nested/", 0750)
	test.DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")
	test.DumpFile(t, fmt.Sprintf("%s/nested/additionals.yaml", tmpdir), "info: data")

	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data", "info": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	os.RemoveAll(tmpdir)
}

func TestParseWithGarbage(t *testing.T) {
	tmpdir := test.MkTempDir(t)
	test.DumpFile(t, fmt.Sprintf("%s/defaults.yml", tmpdir), "test: data")
	os.Mkdir(tmpdir + "/garbage.yml", 0750)
	core.Init()
	core.Set("core.path.metadata", tmpdir)

	expected := Metamap{"test": "data"}
	actual, err := Parse()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	os.RemoveAll(tmpdir)
}
