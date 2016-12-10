package operation

import (
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	pid, err := Execute("sleep", []string{"1"})
	if err != nil {
		t.Fail()
	}

	_, errtest := os.FindProcess(pid)

	if errtest != nil {
		t.Fail()
	}
}
