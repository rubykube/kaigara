// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"github.com/mod/kaigara/pkg/metadata"
	"github.com/mod/kaigara/pkg/resource"
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configmap string

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render NAME",
	Short: "Generate a file from a template",
	Long: `Read metadata for rendering resource pipeline

  In local development it's recommended to setup
  default variables into metadata file
  using yaml, toml, json, xml format.

	Example: kaigara render server.conf > /etc/server.conf`,

	Run: func(cmd *cobra.Command, args []string) {
		metadata.Parse()

		if len(args) > 0 {
			resource.Render("resources/"+args[0]+".tmpl", viper.AllSettings())
		} else {
			term.Error("Error: no template given")
		}
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)
	cobra.OnInitialize(metadata.SetFile)
	// renderCmd.Flags().StringVar(&metaFile,
	// "metafile", "", "Change the metafile path")
}
