package metadata

import (
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/mod/kaigara/pkg/core"
	"github.com/mod/kaigara/pkg/term"
)

func Parse() map[string]interface{} {
	metapath := core.Get("core.path.metadata")
	metafiles, err := ioutil.ReadDir(metapath)
	metadata := map[string]interface{}{}

	if err != nil {
		term.Error(fmt.Sprintf("Error opening path %s", metapath))
		term.Error(err.Error())
	}

	for _, metafile := range metafiles {
		file := fmt.Sprintf("%s/%s", metapath, metafile.Name())
		term.Say("Found Metadata: " + file)
		err, data := LoadYaml(file)
		if err != nil {
			term.Error(err.Error())			
		}
		for k, v := range data {
			metadata[k] = v
		}
	}
	return metadata
}

func LoadYaml(filename string) (error, map[string]interface{}) {
	var err error
	data := map[string]interface{}{}

	fileinfo, err := os.Stat(filename)

	if err != nil {
		return err, nil
	}

	filesize := fileinfo.Size()

	fp, err := os.Open(filename)

	if err != nil {
		return err, nil
	}

	defer fp.Close()

	buf := make([]byte, filesize)
	fp.Read(buf)

	err = yaml.Unmarshal(buf, &data)

	if err != nil {
		return err, nil
	}

	return nil, data
}
