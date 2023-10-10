package match

import "cbtennis/internal/scoring"

type MatchScoreCountControl struct {
	*scoring.ScoreCountControl
}

func newMatchScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &MatchScoreCountControl{ScoreCountControl: scc}
}

func (c MatchScoreCountControl) IsDone(valueA, valueB int) bool {
	if c.UpdateHandler != nil {
		return c.IsDoneHandler(c.MaxValue(), c.HasToConfirm(), valueA, valueB)
	}
	return false
}

func (c MatchScoreCountControl) UpdateScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	sscc := scc.(*MatchScoreCountControl)

	if c.UpdateHandler != nil {
		c.UpdateHandler(sscc, valueA, valueB)
	}
}
