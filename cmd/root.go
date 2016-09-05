// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
  Use:   "kaigara",
  Short: "The devops swiss-army knife",
  Long: `Kaigara is a lightweight provisioning system for unix

  By embeding operations and resources into your application containers
  kaigara will run all your provisioning scripts before starting the app.`,

  // Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kairc)")
}

// initConfig reads in config file and ENV variables if set.
// with the ability to specify config file via flag
func initConfig() {
  if cfgFile != "" {
    viper.SetConfigFile(cfgFile)
  }

  viper.SetConfigName(".kairc")
  viper.AddConfigPath("$HOME")
  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
