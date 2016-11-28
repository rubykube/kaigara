package resource

import (
	"bytes"
	"fmt"

	"github.com/mod/kaigara/pkg/config"
	"github.com/mod/kaigara/pkg/file"
	"github.com/mod/kaigara/pkg/term"

	"text/template"

	_ "github.com/spf13/viper/remote"
)

func tmplPath() []string {
	tmplPath := []string{"./resources", config.Get("tmpl")}
	return tmplPath
}

func Render(f string, data map[string]interface{}) {
	tmpl := file.Read(f)
	rendered := ParseTemplate(tmpl, data)
	fmt.Println(rendered)
}

func RenderConfigMap(tmpl_file string, cm_file string, data map[string]interface{}) {
	cm_data := file.Read(cm_file)
	tmpl := file.Read(tmpl_file)

	cm := ParseConfigMap(cm_data)

	ApplyConfigMap(cm)

	rendered := ParseTemplate(tmpl, data)
	fmt.Println(rendered)
}

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
