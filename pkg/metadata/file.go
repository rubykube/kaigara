package metadata

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/mod/kaigara/pkg/core"
	"github.com/mod/kaigara/pkg/util"
	"github.com/mod/kaigara/pkg/log"
)

type Metamap map[string]interface{}

func Parse() (Metamap, error) {
	metapath := core.Get("core.path.metadata")
	metafiles, err := util.ReadGlob(metapath + "/*.y*ml")
	metadata := Metamap{}

	if err != nil {
		return nil, err
	}

	for _, metafile := range metafiles {
		log.Debug("Found Metadata: " + metafile)
		data, err := LoadYaml(metafile)
		if err != nil {
			log.Error(err.Error())
		}
		for k, v := range data {
			metadata[k] = v
		}
	}
	return metadata, nil
}

func LoadYaml(filename string) (Metamap, error) {
	var err error
	data := Metamap{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
