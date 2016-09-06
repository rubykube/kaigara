// Copyright © 2016 Dimitry Koval <dkoval@heliostech.fr>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

  "os"
  "os/exec"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command or a binary",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    execute(args[0], args[1:])
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}

func execute(cmd string, args []string) {
  process := exec.Command(cmd, args...)
  process.Stdin = os.Stdin
  process.Stdout = os.Stdout
  process.Stderr = os.Stderr

  err := process.Start()
  if err != nil {
    fmt.Println(err)
    return
  }

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		select {
		case <-sigs:
			process.Process.Kill()
		}
	}()

	err = process.Wait()

  if err != nil {
    fmt.Println(err)
    return
  }
}
