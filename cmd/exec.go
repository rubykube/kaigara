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
	"log"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute an executable in a child process",
	Long: `This command can be used as an entrypoint, it will
  search in $PATH for the executable and run it in a child
  process`,
	Run: func(cmd *cobra.Command, args []string) {
		term.Say("exec called")
		operation.RollUp()
		if len(args) > 0 {
			operation.Execute(args[0], args[1:])
		} else {
			log.Fatal("Missing application to start")
		}
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
