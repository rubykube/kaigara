package term

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSay(t *testing.T) {
	actual := Say("kaigara")
	expected := color.New(color.Bold).SprintFunc()("kaigara")
	assert.Equal(t, expected, actual)
}

func TestWarning(t *testing.T) {
	actual := Warning("kaigara")
	expected := color.New(color.FgYellow, color.Bold).SprintFunc()("Warning: kaigara")
	assert.Equal(t, expected, actual)
}

func TestCreate(t *testing.T) {
	actual := Create("kaigara")
	green := color.New(color.FgGreen).SprintFunc()
	expected := color.New(color.Bold).SprintfFunc()("%s kaigara", green("create"))
	assert.Equal(t, expected, actual)
}
