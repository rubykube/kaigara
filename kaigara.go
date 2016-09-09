// Copyright © 2016 Louis <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"github.com/mod/kaigara/cmd"
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("home", "/opt/kaigara")
	viper.SetDefault("path", "/opt/operations")
	viper.SetDefault("tmpl", "/opt/resources")
	viper.SetDefault("env", "development")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kaigara")
	viper.AddConfigPath("$HOME/.kaigara")

	viper.SetEnvPrefix("kaigara")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		term.Say("No configuration found")
	}
}

func main() {
	setDefaults()
	initConfig()
	cmd.Execute()
}
