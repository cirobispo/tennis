package tiebreak

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
	"cbtennis/internal/turning"
)

type TieBreakScoreCountControl struct {
	*scoring.ScoreCountControl

	turn        turning.TurnPosition
	destination gamepoint.GamePointDestination
}

func newTieBreakScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &TieBreakScoreCountControl{ScoreCountControl: scc}
}

func (t *TieBreakScoreCountControl) SetTurn(turn turning.TurnPosition) {
	t.turn = turn
}

func (t *TieBreakScoreCountControl) SetDestination(destination gamepoint.GamePointDestination) {
	t.destination = destination
}

func (t TieBreakScoreCountControl) IsDone(valueA, valueB int) bool {
	if t.UpdateHandler != nil {
		return t.IsDoneHandler(t.MaxValue(), t.HasToConfirm(), valueA, valueB)
	}
	return false
}

func (t TieBreakScoreCountControl) UpdateScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	sscc := scc.(*TieBreakScoreCountControl)

	if t.UpdateHandler != nil {
		t.UpdateHandler(sscc, valueA, valueB)
	}
}
