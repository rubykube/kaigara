package file

import (
	"io/ioutil"
	"os"

	"github.com/mod/kaigara/pkg/term"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFile(filename string) {
	if Exists(filename) {
		term.Error("file '" + filename + "' already exists")
	} else {
		file, err := os.Create(filename)

		if err != nil {
			term.Error(err.Error())
		} else {
			term.Create(filename)
		}

		defer file.Close()
	}
}

func CreateDir(dirname string) {
	if Exists(dirname) {
		term.Error("directory '" + dirname + "' already exists")
	} else {
		err := os.Mkdir(dirname, 0755)

		if err != nil {
			term.Error(err.Error())
		} else {
			term.Create(dirname)
		}
	}
}

func Read(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		term.Error(err.Error())
	}

	return string(file)
}

func Write(filename string, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		term.Error(err.Error())
	}
}
