package set

import (
	"cbtennis/internal/scoring"
	"math"
)

type SetScore struct {
	*scoring.Score
}

func New(scc scoring.ScoringCountControl) scoring.Scoring {
	result := &SetScore{
		Score: scoring.NewCustomGameScore(scoring.STSet, scc),
	}

	scc.PlugToScoring(result.Score)

	return result
}

func NewSetScoreCountControl(maxValue int, hasToConfirm bool) scoring.ScoringCountControl {
	return newSetScoreCountControl(maxValue, hasToConfirm, updateSetScore, isSetFinished)
}

func updateSetScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	sscc := scc.(*SetScoreCountControl)
	score := sscc.game.GetScore()
	gameValueA, gameValueB := score.GetStatus()

	value := valueA
	if gameValueB > gameValueA {
		value = valueB
	}

	*value++
}

func isSetFinished(maxValue int, hasToConfirm bool, valueA, valueB int) bool {
	cp := 0
	if hasToConfirm {
		cp = 1
	}

	simpleWin := (valueA >= maxValue && valueB < maxValue-cp) || (valueB >= maxValue && valueA < maxValue-cp)
	confirmWin := hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > maxValue || valueB > maxValue)

	return simpleWin || confirmWin
}
