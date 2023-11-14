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

func NewScoreControl(maxValue int, hasToConfirm bool) scoring.ScoringCountControl {
	return newScoreControl(maxValue, hasToConfirm, updateTieBreakScore, isTieBreakFinished)
}

func updateTieBreakScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	var value *int
	tcss := scc.(*TieBreakScoreCountControl)

	actualA := valueA
	actualB := valueB
	dinamicDefiant := turning.TPTurnB
	if tcss.defiantSide == turning.TPTurnB {
		actualA = valueB
		actualB = valueA
		dinamicDefiant = turning.TPTurnA
	}

	if tcss.destination == gamepoint.GPDSameSide {
		value = actualA
		if tcss.defiantSide == dinamicDefiant {
			value = actualB
		}
	} else {
		value = actualA
		if tcss.ballCurrentTurn == tcss.ballStartTurn {
			value = actualB
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
