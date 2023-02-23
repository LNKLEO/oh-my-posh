//go:build !windows

package ansi

import "github.com/LNKLEO/oh-my-posh/platform"

func GetAccentColor(env platform.Environment) (*RGB, error) {
	return nil, &platform.NotImplemented{}
}

func (d *DefaultColors) SetAccentColor(env platform.Environment, defaultColor string) {
	if len(defaultColor) == 0 {
		return
	}
	d.accent = &Colors{
		Foreground: string(d.ToColor(defaultColor, false)),
		Background: string(d.ToColor(defaultColor, true)),
	}
}
