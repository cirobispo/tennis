package cbtennis

type MatchScore struct {
	GameScore
}

func newMatchScore(setsToWin int) GameScore {
	maxPoints := newMaxPoints(setsToWin, false, updateGameScore)
	return GameScore{
		valueA:            0,
		valueB:            0,
		typeMode:          STMatch,
		maxPoints:         maxPoints,
		changedScoreEvent: make([]OnScoreChange, 0),
	}
}

func NewSimpleMatch() GameScore {
	return newMatchScore(2)
}

func NewLongMatch() GameScore {
	return newMatchScore(3)
}
