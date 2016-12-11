package resource

import (
	"bytes"
	"fmt"

	"github.com/mod/kaigara/pkg/term"
	"github.com/mod/kaigara/pkg/util"

	"text/template"
)

/*
** Render a template file with given data
*/
func Render(f string, data map[string]interface{}) {
	tmpl := util.Read(f)
	rendered := ParseTemplate(tmpl, data)
	fmt.Print(rendered)
}

/*
** Render a template string with given data
*/
func ParseTemplate(tmpl string, data map[string]interface{}) string {
	var buf bytes.Buffer
	var t template.Template

	_, err := t.Parse(tmpl)
	if err != nil {
		term.Error(err.Error())
	}

	err = t.Execute(&buf, data)
	if err != nil {
		term.Error(err.Error())
	}

	return buf.String()
}
