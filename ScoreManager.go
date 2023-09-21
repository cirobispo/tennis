package tennisstatus

import (
	"fmt"
	"math"
)

type maxPoints struct {
	maxValue     int
	hasToConfirm bool

	updateScore func(maxValue int, AB TurnPosition, valueA, valueB *int)
}

func newMaxPoints(max int, confirm bool, updateScore func(maxValue int, AB TurnPosition, valueA, valueB *int)) maxPoints {
	return maxPoints{maxValue: max, hasToConfirm: confirm, updateScore: updateScore}
}

func (p maxPoints) UpdateScore(AB TurnPosition, valueA, valueB *int) {
	if p.updateScore != nil {
		p.updateScore(p.maxValue, AB, valueA, valueB)
	}
}

func (p maxPoints) NeedsConfirmation() bool {
	return p.hasToConfirm
}

func (p maxPoints) GetToMaxPoint(valueA, valueB int) bool {
	simpleWin := (valueA >= p.maxValue && valueB < p.maxValue-1) || (valueB >= p.maxValue && valueA < p.maxValue-1)
	confirmWin := p.hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > p.maxValue || valueB > p.maxValue)

	return simpleWin || confirmWin
}

type OnChangedScore func()
type OnScoreGame func(valueA, valueB int)

type ScoreType int

const (
	STStandard      ScoreType = 0
	STGame          ScoreType = 1
	STTieBreak      ScoreType = 2
	STSuperTieBreak ScoreType = 3
)

type ScoreManager interface {
	ResetScore()
	UpdateScore(AB TurnPosition)
	AddChangedScoreEvent(changedScoreEvent OnChangedScore)
	AddReachedScoreEvent(scoreGameEvent OnScoreGame)
}

type ScoreData interface {
	ScoreType() ScoreType
	GetScoreValues(AB TurnPosition) (int, int)
}

type ScoreDataWrapper struct {
	valueA, valueB string
}

func NewScoreDataWrapper(AB TurnPosition, scoreData ScoreData) ScoreDataWrapper {
	vA, vB := scoreData.GetScoreValues(AB)

	scoreToText := func(value int) string {
		if scoreData.ScoreType() == STGame {
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

type GameScore struct {
	valueA            int
	valueB            int
	typeMode          ScoreType
	maxPoints         maxPoints
	reachedScoreEvent []OnScoreGame
	changedScoreEvent []OnChangedScore
}

func updateGameScore(maxValue int, AB TurnPosition, valueA, valueB *int) {
	if AB == TPEven && *valueB == maxValue {
		*valueB -= 1
	} else if AB == TPOdd && *valueA == maxValue {
		*valueA -= 1
	} else {
		value := valueA
		if AB == TPOdd {
			value = valueB
		}
		*value++
	}
}

func updateSetScore(maxValue int, AB TurnPosition, valueA, valueB *int) {
	value := valueA
	if AB == TPOdd {
		value = valueB
	}
	*value++
}

func NewGameScore() GameScore {
	maxPoints := newMaxPoints(4, true, updateGameScore)
	return GameScore{valueA: 0, valueB: 0, typeMode: STGame, maxPoints: maxPoints, changedScoreEvent: make([]OnChangedScore, 0)}
}

type SetScore struct {
	GameScore
}

func NewSetScore() GameScore {
	return NewCustomSetScore(STGame, 6, true)
}

func NewCustomSetScore(typeMode ScoreType, gamesToWin int, confirmGame bool) GameScore {
	maxPoints := newMaxPoints(gamesToWin, confirmGame, updateSetScore)
	return GameScore{valueA: 0, valueB: 0, typeMode: typeMode, maxPoints: maxPoints, changedScoreEvent: make([]OnChangedScore, 0)}
}

type MatchScore struct {
	GameScore
}

func NewMatchScore(typeMode ScoreType, setsToWin int) GameScore {
	maxPoints := newMaxPoints(setsToWin, false, updateGameScore)
	return GameScore{valueA: 0, valueB: 0, typeMode: STStandard, maxPoints: maxPoints, changedScoreEvent: make([]OnChangedScore, 0)}
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

func (s *GameScore) AddChangedScoreEvent(changedScoreEvent OnChangedScore) {
	s.changedScoreEvent = append(s.changedScoreEvent, changedScoreEvent)
}

func (s *GameScore) ResetScore() {
	s.valueA = 0
	s.valueB = 0
	s.executeChangedScoreEvent()
}

func (s *GameScore) UpdateScore(AB TurnPosition) {
	s.maxPoints.UpdateScore(AB, &s.valueA, &s.valueB)
	s.executeChangedScoreEvent()

	if s.maxPoints.GetToMaxPoint(s.valueA, s.valueB) {
		s.executeScoreGameEvent(s.valueA, s.valueB)
	}
}

func (s GameScore) GetScoreValues(AB TurnPosition) (int, int) {
	if AB == TPOdd {
		return s.valueB, s.valueA
	}

	return s.valueA, s.valueB
}
