package day1

import (
	"github.com/sesquipedalian-dev/advent2019/pkg/core"
)

func NewDay1() core.Day {
	return core.Day{
		ID: "1",
		Challenges: map[string]core.Challenge{
			"1": core.Challenge{
				ID:      "1",
				Handler: core.ChallengeHandlerWrapped(challenge1Handler),
			},
			"2": core.Challenge{
				ID:      "2",
				Handler: core.ChallengeHandlerWrapped(challenge2Handler),
			},
		},
	}
}
