package app

import (
	"AdventOfCode/pkg/app"
	"fmt"
	"math/rand"
	"time"
)

type Exec20221 app.ChallengeMsg

func (e Exec20221) Exec(sub *chan app.ChallengeMsg) {
	for i := 1; i <= 11; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)+100))
		app.TeaLog(sub, fmt.Sprintf("Challenge %d", i))
	}
}
