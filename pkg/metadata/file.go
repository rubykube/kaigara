package metadata

import (
	"io/ioutil"

	"github.com/mod/kaigara/pkg/config"
	"github.com/mod/kaigara/pkg/term"
)

func Parse() {
	metapath := config.Get("core.path.metadata")
	metafiles, err := ioutil.ReadDir(metapath)

	if err != nil {
		term.Error(err.Error())
	}

	for _, metafile := range metafiles {
		term.Say("Found Metadata: " + metafile.Name())
	}
}
