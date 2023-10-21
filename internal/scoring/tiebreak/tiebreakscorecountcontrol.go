package tiebreak

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
)

type TieBreakScoreCountControl struct {
	*scoring.ScoreCountControl

	ballTurn    turning.TurnPosition
	serveTurn   turning.TurnPosition
	destination gamepoint.GamePointDestination
}

func newTieBreakScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &TieBreakScoreCountControl{ScoreCountControl: scc}
}

func (t *TieBreakScoreCountControl) SetBallTurn(turn turning.TurnPosition) {
	t.ballTurn = turn
}

func (t *TieBreakScoreCountControl) SetServeTurn(turn turning.TurnPosition) {
	t.serveTurn = turn
}

func (t *TieBreakScoreCountControl) SetDestination(destination gamepoint.GamePointDestination) {
	t.destination = destination
}
