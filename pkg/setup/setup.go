package setup

import (
	"github.com/sesquipedalian-dev/advent2019/pkg/core"
	"github.com/sesquipedalian-dev/advent2019/pkg/days/day1"
)

func ConfigureDays() *core.Days {
	days := []core.Day{
		day1.NewDay1(),
	}
	daysMap := map[string]core.Day{}
	for _, day := range days {
		daysMap[day.ID] = day
	}

	return &core.Days{Days: daysMap}
}
