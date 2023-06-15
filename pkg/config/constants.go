package config

import challenges "AdventOfCode/pkg/app/challenges"

const (
	SelectedColor = "#009900"
	ErrorColor    = "#FF5263"
)

var Years = "years"
var Days = "days"
var Challenges = map[string]Challenge{
	"2022-1": challenges.Exec20221{},
}
