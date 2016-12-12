package util

import (
	"io/ioutil"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	file, _ := ioutil.TempFile(".", "file.tmp")
	ex := Exists(file.Name())

	os.Remove(file.Name())

	assert.NotNil(t, ex)
}

func TestReadGlob(t *testing.T) {
	metadir, _ := ioutil.TempDir("", "metadata-")
	CreateFile(metadir + "/values.yml")
	CreateFile(metadir + "/meta.yaml")
	expected := []string{metadir + "/meta.yaml", metadir + "/values.yml"}
	actual, err := ReadGlob(metadir + "/*.y*ml")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	os.RemoveAll(metadir)
}

func TestCreateFile(t *testing.T) {
	filename := "test_file.txt"
	CreateFile(filename)

	_, err := os.Stat(filename)

	assert.Nil(t, err, "did not create a file")

	os.Remove(filename)
}

func TestCreateDir(t *testing.T) {
	dirname := "test_dir"
	CreateDir(dirname)

	_, err := os.Stat(dirname)

	assert.Nil(t, err, "did not create a directory")

	os.Remove(dirname)
}
