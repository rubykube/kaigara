// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kaigara",
	Short: "The devops swiss-army knife",
	Long: `Kaigara is a lightweight provisioning system for unix

  By embeding operations and resources into your application containers
  kaigara will run all your provisioning scripts before starting the app.`,

	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kairc)")
	RootCmd.PersistentFlags().Bool("color", false, "enable colorized output")
	viper.BindPFlag("core.color", RootCmd.PersistentFlags().Lookup("color"))
}

// initConfig reads in config file and ENV variables if set.
// with the ability to specify config file via flag
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".kairc")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		term.Warning("No configuration found")
	} else {
		term.Say("Using config file: " + viper.ConfigFileUsed())
	}
}
