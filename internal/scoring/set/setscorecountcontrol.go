package set

import (
	"cbtennis/internal/game"
	"cbtennis/internal/scoring"
)

type SetScoreCountControl struct {
	*scoring.ScoreCountControl

	game game.GameManager
}

func newSetScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &SetScoreCountControl{ScoreCountControl: scc}
}

func (s *SetScoreCountControl) SetCurrentGame(game game.GameManager) {
	s.game = game
}

func (c SetScoreCountControl) IsDone(valueA, valueB int) bool {
	if c.UpdateHandler != nil {
		return c.IsDoneHandler(c.MaxValue(), c.HasToConfirm(), valueA, valueB)
	}
	return false
}

func (c SetScoreCountControl) UpdateScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	sscc := scc.(*SetScoreCountControl)

	if c.UpdateHandler != nil {
		c.UpdateHandler(sscc, valueA, valueB)
	}
}
