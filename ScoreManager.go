package cbtennis

import (
	"fmt"
)

type OnScoreChange func()
type OnScoreGame func(valueA, valueB int)

type ScoreType int

const (
	STCustom        ScoreType = 0
	STGame          ScoreType = 1
	STSet           ScoreType = 2
	STMatch         ScoreType = 3
	STTieBreak      ScoreType = 4
	STSuperTieBreak ScoreType = 5
)

type ScoreManager interface {
	ResetScore()
	UpdateScore(AB TurnPosition, pointDestination GamePointDestination)
	AddChangedScoreEvent(scoreChangeEvent OnScoreChange)
	AddReachedScoreEvent(scoreGameEvent OnScoreGame)
}

type ScoreDataWrapper struct {
	valueA, valueB string
}

func NewScoreDataWrapper(AB TurnPosition, scoreType ScoreType, vA, vB int) ScoreDataWrapper {
	scoreToText := func(value int) string {
		if scoreType == STGame {
			if value == 0 {
				return "Love"
			} else if value >= 1 && value <= 2 {
				return fmt.Sprint(value * 15)
			} else if value == 3 {
				return "40"
			} else if value == 4 {
				if (vA == 4 && vB == 3) || (vA == 3 && vB == 4) {
					return "Ad"
				}
				return "Game"
			} else if value == 5 {
				return "Game"
			}
		}

		return fmt.Sprint(value)
	}

	if vA == 0 && vB == 0 {
		return ScoreDataWrapper{valueA: "0", valueB: "0"}
	}

	return ScoreDataWrapper{valueA: scoreToText(vA), valueB: scoreToText(vB)}
}

func (w ScoreDataWrapper) GetValueA() string {
	return w.valueA
}

func (w ScoreDataWrapper) GetValueB() string {
	return w.valueB
}
