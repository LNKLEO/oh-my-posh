package segments

import (
	"github.com/LNKLEO/OMP/src/platform"
	"github.com/LNKLEO/OMP/src/properties"
)

type Text struct {
	props properties.Properties
	env   platform.Environment
}

func (t *Text) Template() string {
	return "  "
}

func (t *Text) Enabled() bool {
	return true
}

func (t *Text) Init(props properties.Properties, env platform.Environment) {
	t.props = props
	t.env = env
}
