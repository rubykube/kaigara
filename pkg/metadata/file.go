package metadata

import (
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/mod/kaigara/pkg/core"
	"github.com/mod/kaigara/pkg/term"
)

func Parse() (map[string]interface{}, error) {
	metapath := core.Get("core.path.metadata")
	metafiles, err := ioutil.ReadDir(metapath)
	metadata := map[string]interface{}{}

	if err != nil {
		return nil, err
	}

	for _, metafile := range metafiles {
		file := fmt.Sprintf("%s/%s", metapath, metafile.Name())
		term.Say("Found Metadata: " + file)
		data, err := LoadYaml(file)
		if err != nil {
			return nil, err
		}
		for k, v := range data {
			metadata[k] = v
		}
	}
	return metadata, nil
}

func LoadYaml(filename string) (map[string]interface{}, error) {
	var err error
	data := map[string]interface{}{}

	fileinfo, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()

	fp, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	buf := make([]byte, filesize)
	fp.Read(buf)

	err = yaml.Unmarshal(buf, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
