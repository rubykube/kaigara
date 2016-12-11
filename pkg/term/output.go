package term

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Warning(msg string) string {
	SetConfig()
	warningColor := color.New(color.FgYellow, color.Bold)
	print := warningColor.SprintfFunc()("Warning: %s", msg)

	fmt.Fprintln(os.Stderr, print)

	return print
}

func Error(msg string) string {
	SetConfig()
	red := color.New(color.FgRed, color.Bold)

	print := red.SprintfFunc()("Error: %s", msg)
	fmt.Fprintln(os.Stderr, print)

	return print
}

func Say(msg string) string {
	SetConfig()
	white := color.New(color.Bold)
	print := white.SprintFunc()(msg)

	fmt.Println(print)

	return print
}

func Create(name string) string {
	SetConfig()
	bold := color.New(color.Bold)
	green := color.New(color.FgGreen).SprintFunc()

	print := bold.SprintFunc()(green("create"), fmt.Sprintf(" %s", name))
	fmt.Println(print)

	return print
}

func SetConfig() {
	color.NoColor = false
}
