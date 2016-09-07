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

package app

import (
	"fmt"

	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func Execute(cmd string, args []string) {
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
