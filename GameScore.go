package cbtennis

type GameScore struct {
	valueA            int
	valueB            int
	typeMode          ScoreType
	maxPoints         maxPoints
	reachedScoreEvent []OnScoreGame
	changedScoreEvent []OnScoreChange
}

func updateGameScore(maxValue int, AB TurnPosition, pointDestination GamePointDestination, valueA, valueB *int) {
	allSameSide := (pointDestination == GPDSameSide && AB == TPEven) ||
		(pointDestination == GPDOpositeSide && AB == TPOdd)

	var increment = 1
	var value *int
	if allSameSide {
		value = valueA
		if *valueB > maxValue {
			increment = -1
			value = valueB
		}
	} else {
		value = valueB
		if *valueA > maxValue {
			increment = -1
			value = valueA
		}
	}

	*value += increment
}

func newCustomGameScore(typeMode ScoreType, pointsNeeded int, needsConfirmation bool) GameScore {
	maxPoints := newMaxPoints(pointsNeeded, needsConfirmation, updateGameScore)
	return GameScore{
		valueA:            0,
		valueB:            0,
		typeMode:          typeMode,
		maxPoints:         maxPoints,
		changedScoreEvent: make([]OnScoreChange, 0),
	}
}

func NewGameScore() GameScore {
	maxPoints := newMaxPoints(4, true, updateGameScore)
	return GameScore{
		valueA:            0,
		valueB:            0,
		typeMode:          STGame,
		maxPoints:         maxPoints,
		changedScoreEvent: make([]OnScoreChange, 0),
	}
}

func NewTieBrakeGame() GameScore {
	return newCustomGameScore(STTieBreak, 7, true)
}

func NewSuperTieBrakeGame() GameScore {
	return newCustomGameScore(STTieBreak, 10, true)
}

func (s GameScore) executeChangedScoreEvent() {
	for i := 0; i < len(s.changedScoreEvent); i++ {
		evt := s.changedScoreEvent[i]
		evt()
	}
}

func (s GameScore) executeScoreGameEvent(valueA, valueB int) {
	for i := 0; i < len(s.reachedScoreEvent); i++ {
		evt := s.reachedScoreEvent[i]
		evt(valueA, valueB)
	}
}

func (s GameScore) ScoreType() ScoreType {
	return s.typeMode
}

func (s *GameScore) AddReachedScoreEvent(scoreGameEvent OnScoreGame) {
	s.reachedScoreEvent = append(s.reachedScoreEvent, scoreGameEvent)
}

func (s *GameScore) AddChangedScoreEvent(changedScoreEvent OnScoreChange) {
	s.changedScoreEvent = append(s.changedScoreEvent, changedScoreEvent)
}

func (s *GameScore) ResetScore() {
	s.valueA = 0
	s.valueB = 0
	s.executeChangedScoreEvent()
}

func (s *GameScore) UpdateScore(AB TurnPosition, pointDestination GamePointDestination) {
	s.maxPoints.UpdateScore(AB, pointDestination, &s.valueA, &s.valueB)
	s.executeChangedScoreEvent()
}

func (s GameScore) GetScoreValues(AB TurnPosition) (int, int) {
	if AB == TPOdd {
		return s.valueB, s.valueA
	}

	return s.valueA, s.valueB
}
