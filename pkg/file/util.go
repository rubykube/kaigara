package file

import "os"

func Exists(filename string) bool {

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
