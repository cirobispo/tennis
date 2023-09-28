package cbtennis

type SetScore struct {
	GameScore
}

func NewSetScore() GameScore {
	return NewCustomSetScore(6, true)
}

func NewCustomSetScore(gamesToWin int, confirmGame bool) GameScore {
	maxPoints := newMaxPoints(gamesToWin, confirmGame, updateSetScore)
	return GameScore{
		valueA:            0,
		valueB:            0,
		typeMode:          STSet,
		maxPoints:         maxPoints,
		changedScoreEvent: make([]OnScoreChange, 0),
	}
}

func updateSetScore(maxValue int, AB TurnPosition, pointDestination GamePointDestination, valueA, valueB *int) {
	value := valueA
	if AB == TPOdd {
		value = valueB
	}
	*value++
}
