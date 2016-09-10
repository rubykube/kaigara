package operation

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestApply(t *testing.T) {
	dir := prepareEnv()
	apply(dir)

	_, err := os.Stat("test1.txt")
	cleanup(dir)

	if err != nil {
		t.Fail()
	}
}

func TestRollUp(t *testing.T) {
	dir := prepareEnv()
	script := []byte("#!/bin/sh\ntouch test2.txt")
	ioutil.WriteFile(dir+"/script2.sh", script, 0755)

	RollUp()

	_, err := os.Stat("test1.txt")
	_, err = os.Stat("test2.txt")
	cleanup(dir)

	if err != nil {
		t.Fail()
	}
}

func prepareEnv() (dir string) {
	os.Mkdir("./operations", 0755)
	dir, _ = filepath.Abs("./operations")
	script := []byte("#!/bin/sh\ntouch test1.txt")

	os.Mkdir(dir, 0755)
	ioutil.WriteFile(dir+"/script1.sh", script, 0755)

	return
}

func cleanup(dir string) {
	os.RemoveAll(dir)
	os.Remove("test1.txt")
	os.Remove("test2.txt")
}
