// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"path"

	"github.com/mod/kaigara/pkg/util"
	"github.com/spf13/cobra"
)

var dir string

// initCmd represents the create command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a Kaigara default docker project",
	Long: `Kaigara can be used as a container entrypoint,
  this command will create a basic skeleton with kaigara
  Dockerfile.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CreateFile(path.Join(dir, "Dockerfile"))
		util.CreateFile(path.Join(dir, "defaults.yml"))
		util.CreateDir(path.Join(dir, "operations/"))
		util.CreateDir(path.Join(dir, "resources/"))
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.Flags().String("image", "debian:latest", "Dockerfile base image")
	initCmd.Flags().StringVar(&dir, "path", ".", "Project directory")
}
