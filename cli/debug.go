package cli

import (
	"fmt"
	"time"

	"github.com/LNKLEO/OMP/build"
	"github.com/LNKLEO/OMP/config"
	"github.com/LNKLEO/OMP/prompt"
	"github.com/LNKLEO/OMP/runtime"
	"github.com/LNKLEO/OMP/template"
	"github.com/LNKLEO/OMP/terminal"

	"github.com/spf13/cobra"
)

// debugCmd represents the prompt command
var debugCmd = createDebugCmd()

func init() {
	RootCmd.AddCommand(debugCmd)
}

func createDebugCmd() *cobra.Command {
	debugCmd := &cobra.Command{
		Use:       "debug [bash|zsh|fish|powershell|pwsh|cmd|nu|tcsh|elvish|xonsh]",
		Short:     "Print the prompt in debug mode",
		Long:      "Print the prompt in debug mode.",
		ValidArgs: supportedShells,
		Args:      NoArgsOrOneValidArg,
		Run: func(cmd *cobra.Command, args []string) {
			startTime := time.Now()

			if len(args) == 0 {
				_ = cmd.Help()
				return
			}

			env := &runtime.Terminal{
				CmdFlags: &runtime.Flags{
					Config: configFlag,
					Debug:  true,
					PWD:    pwd,
					Shell:  args[0],
					Plain:  plain,
				},
			}

			env.Init()
			defer env.Close()

			template.Init(env)

			cfg := config.Load(env)

			// add variables to the environment
			env.Var = cfg.Var

			terminal.Init(args[0])
			terminal.BackgroundColor = cfg.TerminalBackground.ResolveTemplate()
			terminal.Colors = cfg.MakeColors()
			terminal.Plain = plain

			eng := &prompt.Engine{
				Config: cfg,
				Env:    env,
				Plain:  plain,
			}

			fmt.Print(eng.PrintDebug(startTime, build.Version))
		},
	}

	debugCmd.Flags().StringVar(&pwd, "pwd", "", "current working directory")
	debugCmd.Flags().BoolVarP(&plain, "plain", "p", false, "plain text output (no ANSI)")

	// Deprecated flags, should be kept to avoid breaking CLI integration.
	debugCmd.Flags().StringVar(&shellName, "shell", "", "the shell to print for")

	// Hide flags that are deprecated or for internal use only.
	_ = debugCmd.Flags().MarkHidden("shell")

	return debugCmd
}
