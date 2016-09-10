package resource

import (
	"github.com/mod/kaigara/pkg/config"
	"github.com/mod/kaigara/pkg/file"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func tmplPath() []string {
	tmplPath := []string{"./resources", config.Get("tmpl")}
	return tmplPath
}

func Render(tmpl string, data map[string]interface{}) {
	var resourcesPath string

	for _, dir := range tmplPath() {
		if file.Exists(dir) {
			resourcesPath = dir
			break
		}
	}

	fullpath := filepath.Join(resourcesPath, tmpl)
	t, err := template.ParseFiles(fullpath + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
