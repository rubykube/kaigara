package term

import (
	"github.com/fatih/color"
	"testing"
)

func TestSay(t *testing.T) {
	str := Say("kaigara")
	test := color.New(color.Bold).SprintFunc()("kaigara")

	if str != test {
		t.Error()
	}
}

func TestWarning(t *testing.T) {
	str := Warning("kaigara")
	test := color.New(color.Bold, color.FgRed, color.FgYellow).SprintFunc()("Warning: kaigara")

	if str != test {
		t.Error()
	}
}

func TestCreate(t *testing.T) {
	str := Create("kaigara")
	green := color.New(color.FgGreen).SprintFunc()
	test := color.New(color.Bold).SprintfFunc()("%s kaigara", green("create"))

	if str != test {
		t.Error()
	}
}
