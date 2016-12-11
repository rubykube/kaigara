// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"github.com/mod/kaigara/pkg/core"
	"github.com/mod/kaigara/pkg/metadata"
	"github.com/mod/kaigara/pkg/term"
	"github.com/mod/kaigara/pkg/resource"
	"github.com/spf13/cobra"
)

var configmap string

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render TEMPLATE",
	Short: "Generate a file on STDOUT from a TEMPLATE",
	Long: `Read metadata files for rendering resource pipeline

In local development it is recommended to mount default.yml
using -v default.yml:/etc/kaigara/default.yml

Example: kaigara render server.conf > /etc/server.conf`,

	Run: func(cmd *cobra.Command, args []string) {
		core.Init()
		data, err := metadata.Parse()

		if err != nil {
			term.Error(err.Error())
		}

		if len(args) > 0 {
			resource.Render(core.Get("core.path.resources") + "/" + args[0] + ".tmpl", data)
		} else {
			term.Error("No template given")
		}
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)
}
