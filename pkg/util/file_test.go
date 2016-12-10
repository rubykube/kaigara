package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	file, _ := ioutil.TempFile(".", "file.tmp")
	ex := Exists(file.Name())

	os.Remove(file.Name())

	if !ex {
		t.Fail()
	}
}

func TestReadGlob(t *testing.T) {
	metadir, _ := ioutil.TempDir("", "metadata-")
	CreateFile(metadir + "/values.yml")
	CreateFile(metadir + "/meta.yaml")
	files := ReadGlob(metadir + "/*.y*ml")
	for _, file := range files {
		ex := Exists(file)
		if !ex {
			t.Error(ex)
		}
	}
	os.RemoveAll(metadir)
}

func TestCreateFile(t *testing.T) {
	filename := "test_file.txt"
	CreateFile(filename)

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		t.Error("did not create file")
	}

	os.Remove(filename)
}

func TestCreateDir(t *testing.T) {
	dirname := "test_dir"
	CreateDir(dirname)

	_, err := os.Stat(dirname)

	if os.IsNotExist(err) {
		t.Error("did not create dir")
	}

	os.Remove(dirname)
}
