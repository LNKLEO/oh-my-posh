package engine

import (
	"github.com/LNKLEO/oh-my-posh/ansi"
	"github.com/LNKLEO/oh-my-posh/platform"
	"github.com/LNKLEO/oh-my-posh/shell"
)

// New returns a prompt engine initialized with the
// given configuration options, and is ready to print any
// of the prompt components.
func New(flags *platform.Flags) *Engine {
	env := &platform.Shell{
		CmdFlags: flags,
	}

	env.Init()
	cfg := LoadConfig(env)
	env.Var = cfg.Var

	ansiWriter := &ansi.Writer{
		TerminalBackground: shell.ConsoleBackgroundColor(env, cfg.TerminalBackground),
		AnsiColors:         cfg.MakeColors(),
		Plain:              flags.Plain,
	}
	ansiWriter.Init(env.Shell())

	eng := &Engine{
		Config: cfg,
		Env:    env,
		Writer: ansiWriter,
		Plain:  flags.Plain,
	}

	return eng
}
