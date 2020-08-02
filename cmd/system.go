/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/splunk-golang/internal"
	splunk "github.com/brittonhayes/splunk-golang/pkg"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	confirm bool
)

func init() {
	//Commands
	rootCmd.AddCommand(SystemCmd)
	SystemCmd.AddCommand(restartCmd)
	SystemCmd.AddCommand(controlsCmd)

	//Flags
	restartCmd.Flags().BoolVarP(&confirm, "confirm", "c", false, "Confirm you would like to restart.")

	//Colors
	au = aurora.NewAurora(*colors)
}

// SystemCmd represents the system command
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Perform operations on the Splunk server.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
	},
}

// restartCmd is used to restart the Splunk instance
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Splunk instance",
	Run: func(cmd *cobra.Command, args []string) {

		validate := func(input string) error {
			if len(input) < 1 {
				return errors.New("provide an input to confirm")
			}
			return nil
		}

		if !confirm {
			confirmPrompt := promptui.Prompt{
				Label:    "Are you sure you'd like to restart Splunk? [Y/n]",
				Validate: validate,
				Default:  "Y",
			}

			result, err := confirmPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}

			if result == "Y" {
				msg := restartSplunk()
				fmt.Println(msg)
			} else {
				os.Exit(0)
			}

		}

	},
}

// controlsCmd is used to restart the Splunk instance
var controlsCmd = &cobra.Command{
	Use:   "controls [name]",
	Short: "Lists actions that can be performed at this endpoint.",
	Long: `
	Function: Lists actions that can be performed at this endpoint. 
	Format: JSON
	Tip: Pipe into jq for prettified output`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
		fmt.Println(inspectControl(args[0]))
	},
}

func restartSplunk() string {
	conn := splunk.Connection{
		Username: viper.GetString("SPLUNK_USERNAME"),
		Password: viper.GetString("SPLUNK_PASSWORD"),
		BaseURL:  viper.GetString("SPLUNK_URL"),
	}
	response, err := conn.RestartServer()
	if err != nil {
		log.Fatal(au.Red("Couldn't restart Splunk. Ensure your configuration is correct."))
	}

	return response
}

func inspectControl(endpoint string) string {
	conn := splunk.Connection{
		Username: viper.GetString("SPLUNK_USERNAME"),
		Password: viper.GetString("SPLUNK_PASSWORD"),
		BaseURL:  viper.GetString("SPLUNK_URL"),
	}
	response, err := conn.InspectControl(endpoint)
	if err != nil {
		log.Fatal(au.Red("Couldn't list actions for this endpoint. Check your input and try again."))
	}
	return response
}