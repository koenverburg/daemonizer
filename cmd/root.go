/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/koenverburg/daemonizer/utils"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "daemonizer",
	Short: "A tool to start processes in a tmux session",
	Long: "A tool to start processes in a tmux session",
	Run: func(cmd *cobra.Command, args []string) {
    settings := viper.AllSettings()
    utils.CreateTmuxWindows("dotfiles", settings)
    // fmt.Println(settings)
    // utils.CreateTmuxSession()

    // if commandSet == nil {
    //   panic(fmt.Sprintf("No commands found for %s", args[0]))
    // }

    // for _, command := range commandSet {
    //   utils.RunCommand(command)
    // }
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.daemonizer.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		config, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in config directory with name ".daemonizer" (without extension).
		viper.AddConfigPath(config)
		viper.SetConfigName(".daemonizer")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
