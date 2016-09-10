package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func ExistsTest(t *testing.T) {
	file, _ := ioutil.TempFile(".", "file.tmp")
	ex := Exists(file.Name())

	if !ex {
		t.Fail()
	}

	os.Remove(file.Name())
}
