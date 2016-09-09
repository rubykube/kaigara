// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Inventory struct {
	Material string
	Count    uint
}

var metaFile string

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Generate a file from a template",
	Long: `Read metadata for rendering resource pipeline

  In local development it's recommended to setup
  default variables into metadata file
  using yaml, toml, json, xml format.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("render called with %s", args[0])

			data := Inventory{"wool", 17}
			renderTemplate(args[0], &data)
		}
	},
}

func renderTemplate(tmpl string, data *Inventory) {
	t, err := template.ParseFiles("resources/" + tmpl + ".tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(renderCmd)
	cobra.OnInitialize(setMetadata)
	renderCmd.Flags().StringVar(&metaFile, "metafile", "", "Change the metafile path")
}

func setMetadata() {
	if metaFile != "" {
		viper.SetConfigFile(metaFile)
		fmt.Println("setMetafile")
	}

	viper.SetConfigName("metadata")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using metafile:", viper.ConfigFileUsed())
	}
}
