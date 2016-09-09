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
	"log"
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

var Path []string = []string{".", "/opt/kaigara"}

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
			renderTemplate(args[0], viper.AllSettings())
		} else {
			log.Fatal("Error: no template given")
		}
	},
}

func renderTemplate(tmpl string, data map[string]interface{}) {
	var resourcesPath string

	for _, dir := range Path {
		if _, err := os.Stat(dir + "/resources"); err == nil {
			resourcesPath = dir
			break
		}
	}

	t, err := template.ParseFiles(resourcesPath + "/resources/" + tmpl + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(renderCmd)
	cobra.OnInitialize(setMetadata)
	renderCmd.Flags().StringVar(&metaFile, "metafile", "", "Change the metafile path")
}

// TODO: implement KAIGARA_PATH for templates
func setMetadata() {
	if metaFile != "" {
		viper.SetConfigFile(metaFile)
	}

	viper.SetConfigName("metadata")

	for _, dir := range Path {
		if _, err := os.Stat(dir + "/metadata.yml"); err == nil {
			viper.AddConfigPath(dir)
			break
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using metafile:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}
