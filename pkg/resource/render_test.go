package resource

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	data := map[string]interface{}{
		"name":  "kaigara",
		"value": 123,
	}
	tmpl := "{{ .name }} value is {{ .value }}"
	actual, _ := ParseTemplate(tmpl, data)
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

	data := map[string]interface{}{
		"name":  "kaigara",
		"value": 1234,
	}

	actual, _ := Render("test.tmpl", data)
	expected := "kaigara value is 1234"
	assert.Equal(t, expected, actual)
	os.Remove("test.tmpl")
}
