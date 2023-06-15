package app

type ChallengeMsg struct {
	Data string
}

func TeaLog(sub *chan ChallengeMsg, msg string) {
	*sub <- ChallengeMsg{
		Data: msg,
	}
}
