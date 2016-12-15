package operation

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/mod/kaigara/pkg/core"
	"github.com/mod/kaigara/pkg/log"
	"github.com/mod/kaigara/pkg/util"
)

func runOps(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("No operations found in folder: " + path)
	}

	for _, file := range files {
		log.Info("Found operation: " + file.Name())
		_, err := Execute(filepath.Join(path, file.Name()), nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func RollUp() error {
	operationDir := core.Get("core.path.operations")
	if util.Exists("./operations") {
		err := runOps("./operations")
		if err != nil {
			return err
		}
	} else if util.Exists(operationDir) {
		err := runOps(operationDir)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Missing operation folder")
	}
	return nil
}
