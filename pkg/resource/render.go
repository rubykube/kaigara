package resource

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/rubykube/kaigara/pkg/metadata"
	"github.com/rubykube/kaigara/pkg/util"
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
	var tpl *template.Template

	tpl = createTemplate()
	tpl, err := tpl.Parse(tmpl)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func createTemplate() *template.Template {
	fmap := sprig.TxtFuncMap()
	return template.New("base").Funcs(fmap)
}
