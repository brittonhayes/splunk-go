package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completions",
	Long: `To load completions:

	Bash:
	
	$ source <(splunk-go completion bash)
	
	# To load completions for each session, execute once:
	Linux:
		$ splunk-gopletion bash > /etc/bash_completion.d/splusplunk-go
	MacOS:
		$ splunk-gopletion bash > /usr/local/etc/bash_completion.d/splusplunk-go
	
	Zsh:
	
	$ source <(splunk-gopletion zsh)
	
	# To load completions for each session, execute once:
	$ splunk-gopletion zsh > "${fpath[1]}/_splusplunk-go
	`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}
