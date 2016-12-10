package resource

import (
	"testing"
)

var testRenderData = `
test1: 123
test2: kaigara
`

func TestParseTemplate(t *testing.T) {
	data := map[string]string {
	}
	ParseTemplate(testRenderData, data)
	t.Fail()
}
