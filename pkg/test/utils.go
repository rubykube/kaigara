package test

import (
	"testing"
	"os"
	"os/exec"
	"bytes"
	"strings"
)

func MkTempDir(t *testing.T) string {
	cmd := exec.Command("mktemp", "-d")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(out.String())
}

func DumpFile(t *testing.T, path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}

	file.WriteString(content)
	file.Close()
}
