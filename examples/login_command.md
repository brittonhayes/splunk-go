```go
package cmd

import (
	"fmt"
	"log"

	splunk "github.com/brittonhayes/splunk-go/pkg"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(LoginCmd)
	au = aurora.NewAurora(*colors)
}

// LoginCmd is used to authenticate to Splunk
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate to Splunk and return a session token.",
	Run: func(cmd *cobra.Command, args []string) {

		conn := splunk.Connection{
			Username: viper.GetString("SPLUNK_USERNAME"),
			Password: viper.GetString("SPLUNK_PASSWORD"),
			BaseURL:  viper.GetString("SPLUNK_URL"),
		}

		key, err := conn.Login()
		if err != nil {
			log.Fatal(au.Red("Couldn't login to Splunk. Ensure your configuration is correct."))
		}

		fmt.Println("Logged in Successfully. \nSession key:", au.BrightBlue(key.Value))
	},
}
```
