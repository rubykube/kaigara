// Copyright © 2016 Louis Bellet <mod@helios.sh>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type Inventory struct {
	Material string
	Count    uint
}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Generate a file from a template",
	Long: `Read metadata for rendering resource pipeline

  In local development it's recommended to setup
  default variables into metadata file
  using yaml, toml, json, xml format.
  define variable for each runtime environment`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("render called with %s", args[0])

		data := Inventory{"wool", 17}
		renderTemplate(args[0], &data)
	},
}

func renderTemplate(tmpl string, data *Inventory) {
	t, err := template.ParseFiles("resources/" + tmpl + ".tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(renderCmd)
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
