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
	var value *int
	tcss := scc.(*TieBreakScoreCountControl)

	if tcss.destination == gamepoint.GPDSameSide {
		value = valueA
		if tcss.serveTurn == turning.TPTurnB {
			value = valueB
		}
	} else {
		if tcss.ballCurrentTurn == tcss.ballStartTurn {
			if tcss.serveTurn == turning.TPTurnA {
				value = valueB
			} else if tcss.serveTurn == turning.TPTurnB {
				value = valueA
			}
		} else {
			if tcss.ballCurrentTurn != tcss.serveTurn {
				value = valueA
				if tcss.serveTurn == turning.TPTurnB {
					value = valueB
				}
			}
		}
	}

	*value += 1
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
