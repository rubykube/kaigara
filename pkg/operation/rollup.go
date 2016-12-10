package operation

import (
	"io/ioutil"
	"path/filepath"

	"github.com/mod/kaigara/pkg/file"
	"github.com/mod/kaigara/pkg/term"
)

func apply(operationPath string) {
	files, err := ioutil.ReadDir(operationPath)

	if err != nil {
		term.Error(err.Error())
	}

	for _, file := range files {
		term.Say("Found operation: " + file.Name())
		Execute(filepath.Join(operationPath, file.Name()), nil)
	}
}

func RollUp() {
	operationDir := "/opt/provision/operation"
	if file.Exists("./operations") {
		apply("./operations")
	} else if file.Exists(operationDir) {
		apply(operationDir)
	} else {
		term.Error("Missing operation folder")
	}
}
