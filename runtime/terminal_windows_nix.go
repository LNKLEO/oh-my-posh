//go:build !darwin

package runtime

import (
	"time"

	"github.com/LNKLEO/OMP/runtime/battery"
)

func (term *Terminal) BatteryState() (*battery.Info, error) {
	defer term.Trace(time.Now())
	info, err := battery.Get()
	if err != nil {
		term.Error(err)
		return nil, err
	}
	return info, nil
}
