// Copyright © 2016 Louis <mod@helios.sh>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"github.com/mod/kaigara/cmd"
	"github.com/mod/kaigara/pkg/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
