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

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Kaigara default docker project",
	Long: `Kaigara can be used as a container entrypoint,
  this command will create a basic skeleton with kaigara
  Dockerfile.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		fmt.Println("Dockerfile")
		fmt.Println("metadata.yml")
		fmt.Println("operations/")
		fmt.Println("resources/")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	createCmd.Flags().String("image", "debian:latest", "Dockerfile base image")
}
