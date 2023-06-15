package config

import "AdventOfCode/pkg/app"

type Challenge interface {
	Exec(*chan app.ChallengeMsg)
}
