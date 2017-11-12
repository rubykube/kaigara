// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"github.com/mod/kaigara/pkg/operation"
	"github.com/mod/kaigara/pkg/term"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start PROGRAM",
	Short: "Runs a <PROGRAM> after provision",
	Long: `Run all operations and then start the PROGRAM

Operations should be protected against re-executions.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := operation.RollUp()
		if (err != nil) {
			term.Error(err.Error())
			return;
		}

		if len(args) > 0 {
			operation.Execute(args[0], args[1:])
		} else {
			term.Error("Missing application to start")
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
