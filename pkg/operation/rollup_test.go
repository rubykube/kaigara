package operation

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestApply(t *testing.T) {
	dir := prepareEnv()
	defer cleanup(dir)
	runOps(dir)

	_, err := os.Stat("test1.txt")

	if err != nil {
		t.Fail()
	}
}

func TestRollUp(t *testing.T) {
	dir := prepareEnv()
	defer cleanup(dir)
	script := []byte("#!/bin/sh\ntouch test2.txt")
	ioutil.WriteFile(dir+"/script2.sh", script, 0755)

	RollUp()

	_, err := os.Stat("test1.txt")
	_, err = os.Stat("test2.txt")

	if err != nil {
		t.Fail()
	}
}

func prepareEnv() string {
	script := []byte("#!/bin/sh\ntouch test1.txt\n")

	dir, err := filepath.Abs("./operations")
	if err != nil {
		log.Fatal("filepath.Abs failed")
	}
	err = os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
	ioutil.WriteFile(dir+"/script1.sh", script, 0755)

	return dir
}

func cleanup(dir string) {
	os.RemoveAll(dir)
	os.Remove("test1.txt")
	os.Remove("test2.txt")
}
