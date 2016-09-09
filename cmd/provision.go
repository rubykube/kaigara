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
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mod/kaigara/pkg/app"
	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provisioning using operations",
	Long: `Kaigara provision is a powerful system for maintaining container:

  Step 1: Verify the state of operations using KAIGARA_INDEX
  Step 2: Execute all pending operations using exec
  Step 3: Update the index to reflect last operations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provision called")
		runOperations()
	},
}

func runOperations() {
	path_ops := filepath.Join("opt", "kaigara", "operations")
	files, err := ioutil.ReadDir(path_ops)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println("Found operation: " + file.Name())
		app.Execute(filepath.Join(path_ops, file.Name()), nil)
	}
}

func init() {
	RootCmd.AddCommand(provisionCmd)

	provisionCmd.Flags().BoolP("dry", "y", false, "Dry run operations")
	provisionCmd.Flags().BoolP("verbose", "v", false, "Activate verbose")
}
