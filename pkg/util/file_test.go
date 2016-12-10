package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func ExistsTest(t *testing.T) {
	file, _ := ioutil.TempFile(".", "file.tmp")
	ex := Exists(file.Name())

	os.Remove(file.Name())

	if !ex {
		t.Fail()
	}
}

func CreateFileTest(t *testing.T) {
	filename := "test_file.txt"
	CreateFile(filename)

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		t.Error("did not create file")
	}

	os.Remove(filename)
}

func CreateDirTest(t *testing.T) {
	dirname := "test_dir"
	CreateDir(dirname)

	_, err := os.Stat(dirname)

	if os.IsNotExist(err) {
		t.Error("did not create dir")
	}

	os.Remove(dirname)
}
