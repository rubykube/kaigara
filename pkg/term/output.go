package term

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func Warning(msg string) string {
	set_config()
	orange := color.New(color.FgRed, color.FgYellow, color.Bold)
	print := orange.SprintfFunc()("Warning: %s", msg)

	fmt.Println(print)

	return print
}

func Error(msg string) string {
	set_config()
	red := color.New(color.FgRed, color.Bold)

	print := red.SprintfFunc()("Error: %s", msg)
	fmt.Println(print)

	return print
}

func Say(msg string) string {
	set_config()
	white := color.New(color.Bold)
	print := white.SprintFunc()(msg)

	fmt.Println(print)

	return print
}

func Create(name string) string {
	set_config()
	bold := color.New(color.Bold)
	green := color.New(color.FgGreen).SprintFunc()

	print := bold.SprintFunc()(green("create "), name)
	fmt.Println(print)

	return print
}

func set_config() {
	color.NoColor = !viper.GetBool("core.color")
}
