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

	"github.com/jubnzv/go-tmux"
	"github.com/koenverburg/daemonizer/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

type Task struct {
	Root     string   `mapstructure:"root"`
	Commands []string `mapstructure:"commands"`
}

type Tasks map[string]Task

var rootCmd = &cobra.Command{
	Use:   "daemonizer",
	Short: "A tool to start processes in a tmux session",
	Long:  "A tool to start processes in a tmux session",
	Run: func(cmd *cobra.Command, args []string) {
		// namespace := args[0]

		var tasks Tasks
		settings := viper.AllSettings()
		mapstructure.Decode(settings, &tasks)

		namespace := "chk"
		directory := tasks["chk"].Root
		commands := tasks["chk"].Commands

		server, session := utils.CreateServer(namespace)
		sessions := []*tmux.Session{}

		for k, v := range commands {
			sessions = append(sessions, &session)
			utils.AddWindow(server, session, directory, v, k)
		}

		utils.Start(server, sessions)
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
		viper.SetConfigName(".background")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
