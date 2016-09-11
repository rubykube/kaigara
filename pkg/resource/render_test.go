package resource

import (
	"github.com/mod/kaigara/pkg/metadata"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	dir := "./resources"
	os.Mkdir(dir, 0755)

	mdata := []byte("var1: test1\nvar2: test2")
	ioutil.WriteFile("metadata.yml", mdata, 0755)
	metadata.Parse()

	tpl := []byte("{{.var1}} {{.var2}}\n")
	ioutil.WriteFile(dir+"/file.txt.tmpl", tpl, 0755)

	Render("file.txt", viper.AllSettings())

	os.Remove("metadata.yml")
	os.RemoveAll(dir)
}
