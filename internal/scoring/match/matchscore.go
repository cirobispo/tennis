package match

import (
	"cbtennis/internal/scoring"
	"math"
)

type MatchScore struct {
	*scoring.Score
}

func New(scc scoring.ScoringCountControl) scoring.Scoring {
	result := &MatchScore{
		Score: scoring.NewCustomGameScore(scoring.STSet, scc),
	}

	scc.PlugToScoring(result.Score)

	return result
}

func NewMatchScoreCountControl(maxValue int, hasToConfirm bool) scoring.ScoringCountControl {
	return newMatchScoreCountControl(maxValue, hasToConfirm, updateSetScore, isSetFinished)
}

func updateSetScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
}

func isSetFinished(maxValue int, hasToConfirm bool, valueA, valueB int) bool {
	simpleWin := (valueA >= maxValue && valueB < maxValue-1) || (valueB >= maxValue && valueA < maxValue-1)
	confirmWin := hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > maxValue || valueB > maxValue)

	return simpleWin || confirmWin
}
