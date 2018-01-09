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

	"github.com/rubykube/kaigara/pkg/core"
	"github.com/rubykube/kaigara/pkg/log"
	"github.com/rubykube/kaigara/pkg/metadata"
	"github.com/rubykube/kaigara/pkg/resource"
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

		if len(args) > 0 {
			data, err := metadata.Parse()
			template := fmt.Sprintf("%s/%s.tmpl", core.Get("core.path.resources"), args[0])

			if err != nil {
				log.Fatal(err.Error())
			}

			output, err := resource.Render(template, data)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Print(output)
		} else {
			log.Error("No template given")
		}
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)
}
