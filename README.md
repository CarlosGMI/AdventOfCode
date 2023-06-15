This is a simple TUI application for listing and executing all the [Advent of Code](https://adventofcode.com/) challenges. The application makes use of [Bubbletea](https://github.com/charmbracelet/bubbletea) for all the lists and options in the application.

## Requirements

You will need to have Golang installed and a terminal in order to run this application.

## How to use

1. Clone the repository into a directory of your preference

```bash
git clone https://github.com/CarlosGMI/AdventOfCode
```

2. Run `go mod tidy`
3. Run `go run ./cmd/main.go`

## Adding challenges

If you want to add more Advent of Code challenges from the ones in this repo you need to:

-   Create a new `README.md` file inside the `pkg/app/instructions` directory making sure you follow the `$YEAR/$DAY` subdirectory structure. For example, if you want to add the 2024 Day 1 challenge you will need to create a new `2024` folder and inside you will create a `1` folder.
    -   If you don't do this, the application won't be able to display the challenge instructions when you're running the app.
-   Create a new `.go` file inside the `pkg/app/challenges` directory. This will be the file that will contain the code for the challenge.
    -   This file needs to implement the `Challenge` interface in the `types.go` file.
-   Add a new entry in the `Challenges` map of the `constants.go` file. The key of the new entry should follow the `$YEAR-$DAY` naming convention. For example, if you want to add the 2024 Day 1 challenge you can add the entry as:

```go
var Challenges = map[string]Challenge{
    "2022-1": challenges.Exec20221{},
    "2024-1": challenges.Exec20241{}, // <-- new challenge
}
```
