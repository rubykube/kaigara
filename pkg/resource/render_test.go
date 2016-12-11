package resource

import (
	"os"
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	stdoutBackup := os.Stdout
	os.Stdout, _ = os.Create("test.output")
	f()
	os.Stdout.Close()
	os.Stdout = stdoutBackup
	file, _ := ioutil.ReadFile("test.output")
	os.Remove("test.output")
	return string(file)
}

func TestParseTemplate(t *testing.T) {
	data := map[string]interface{} {
		"name": "kaigara",
		"value": 123,
	}
	tmpl := "{{ .name }} value is {{ .value }}"
	actual := ParseTemplate(tmpl, data)
	expected := "kaigara value is 123"
	assert.Equal(t, expected, actual)
}

func TestRender(t *testing.T) {
	file, err := os.Create("test.tmpl")
	if err != nil {
		t.Fatal(err)
	}

	file.WriteString("{{ .name }} value is {{ .value }}")
	file.Close()

	data := map[string]interface{} {
		"name": "kaigara",
		"value": 1234,
	}

	actual := captureOutput(func() {
		Render("test.tmpl", data)
	})
	expected := "kaigara value is 1234"
	assert.Equal(t, expected, actual)
	os.Remove("test.tmpl")
}
