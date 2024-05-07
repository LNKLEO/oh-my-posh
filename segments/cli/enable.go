package cli

import (
	"fmt"


	"github.com/spf13/cobra"
)

// getCmd represents the get command
var enableCmd = &cobra.Command{
	Use:       fmt.Sprintf(toggleUse, "enable"),
	Short:     "Enable a feature",
	Long:      fmt.Sprintf(toggleLong, "Enable", "enable"),
	ValidArgs: toggleArgs,
	Args:      NoArgsOrOneValidArg,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		toggleFeature(cmd, args[0], true)
	},
}

func init() {
	RootCmd.AddCommand(enableCmd)
}

func toggleFeature(cmd *cobra.Command, feature string, enable bool) {
	env := &platform.Shell{
		CmdFlags: &platform.Flags{
			Shell: shellName,
		},
	}
	env.Init()
	defer env.Close()
	switch feature {
	default:
		_ = cmd.Help()
	}
}
