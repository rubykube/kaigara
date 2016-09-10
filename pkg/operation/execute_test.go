package operation

import (
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	pid := Execute("sleep", []string{"1"})
	_, err := os.FindProcess(pid)

	if err != nil {
		t.Fail()
	}
}
