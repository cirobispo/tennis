package tiebreak

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

type TieBreakScore struct {
	*scoring.Score
}

func New(scc scoring.ScoringCountControl) scoring.Scoring {
	result := &TieBreakScore{
		Score: scoring.NewCustomGameScore(scoring.STSet, scc),
	}

	scc.PlugToScoring(result.Score)

	return result
}

func NewTieBreakScoreCountControl(maxValue int, hasToConfirm bool) scoring.ScoringCountControl {
	return newTieBreakScoreCountControl(maxValue, hasToConfirm, updateTieBreakScore, isTieBreakFinished)
}

func updateTieBreakScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	tcss := scc.(*TieBreakScoreCountControl)

	var increment = 1
	var value *int

	ballAtMySide := (tcss.serveTurn == turning.TPTurnA && tcss.ballTurn == turning.TPTurnA) ||
		(tcss.serveTurn == turning.TPTurnB && tcss.ballTurn == turning.TPTurnB)

	pointSameSide := (ballAtMySide && tcss.destination == gamepoint.GPDSameSide)

	if *valueB > scc.MaxValue() || *valueA > scc.MaxValue() {
		increment = -1
	}

	if pointSameSide && tcss.serveTurn == turning.TPTurnA {
		value = valueA
		if *valueB > scc.MaxValue() {
			value = valueB
		}
	} else if pointSameSide && tcss.serveTurn == turning.TPTurnB {
		value = valueB
		if *valueA > scc.MaxValue() {
			value = valueA
		}
	} else {
		if tcss.destination == gamepoint.GPDOpositeSide {
			if tcss.serveTurn == turning.TPTurnA {
				value = valueB
				if *valueA > scc.MaxValue() {
					value = valueA
				}
			} else {
				value = valueA
				if *valueB > scc.MaxValue() {
					value = valueB
				}
			}
		} else {
			if tcss.serveTurn == turning.TPTurnA {
				value = valueA
				if *valueB > scc.MaxValue() {
					value = valueB
				}
			} else {
				value = valueB
				if *valueA > scc.MaxValue() {
					value = valueA
				}
			}
		}
	}
	*value += increment
}

func isTieBreakFinished(maxValue int, hasToConfirm bool, valueA, valueB int) bool {
	cp := 0
	if hasToConfirm {
		cp = 1
	}

	simpleWin := (valueA >= maxValue && valueB < maxValue-cp) || (valueB >= maxValue && valueA < maxValue-cp)
	confirmWin := hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > maxValue || valueB > maxValue)

	return simpleWin || confirmWin
}
