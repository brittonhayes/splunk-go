package internal

import (
	"os"

	"github.com/spf13/cobra"
)

// Help checks for input. Defaults to run help command.
func Help(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Help()
		os.Exit(0)
	}
}
