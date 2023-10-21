package match

import "cbtennis/internal/scoring"

type MatchScoreCountControl struct {
	*scoring.ScoreCountControl
}

func newMatchScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &MatchScoreCountControl{ScoreCountControl: scc}
}
