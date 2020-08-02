```go
package cmd

import (
	"fmt"
	"log"

	"github.com/brittonhayes/splunk-golang/internal"
	splunk "github.com/brittonhayes/splunk-golang/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(appsCmd)
	appsCmd.AddCommand(inspectCmd)
}

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Create, inspect, update, or remove Splunk applications.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
	},
}

// inspectCmd lists application's details and properties
var inspectCmd = &cobra.Command{
	Use:   "inspect [name]",
	Short: "Inspects details and properties of the queried application.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Help(cmd, args)
		fmt.Println(appInspect(args[0]))
	},
}

func appInspect(path string) string {
	conn := splunk.Connection{
		Username: viper.GetString("SPLUNK_USERNAME"),
		Password: viper.GetString("SPLUNK_PASSWORD"),
		BaseURL:  viper.GetString("SPLUNK_URL"),
	}
	response, err := conn.AppInspect(path)
	if err != nil {
		log.Fatal(au.Red("Couldn't list Splunk apps. Ensure your configuration is correct."))
	}

	return response
}
```
