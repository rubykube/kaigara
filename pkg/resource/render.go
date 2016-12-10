package resource

import (
	"bytes"
	"text/template"

	"github.com/mod/kaigara/pkg/metadata"
	"github.com/mod/kaigara/pkg/util"
)

/*
** Render a template file with given data
 */
func Render(f string, data metadata.Metamap) (string, error) {
	tmpl := util.Read(f)
	output, err := ParseTemplate(tmpl, data)
	if err != nil {
		return "", err
	}
	return output, nil
}

/*
** Render a template string with given data
 */
func ParseTemplate(tmpl string, data metadata.Metamap) (string, error) {
	var buf bytes.Buffer
	var tpl template.Template

	_, err := tpl.Parse(tmpl)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
