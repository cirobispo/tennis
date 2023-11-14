package tiebreak

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
)

type TieBreakScoreCountControl struct {
	*scoring.ScoreCountControl

	ballStartTurn   turning.TurnPosition
	ballCurrentTurn turning.TurnPosition
	defiantSide     turning.TurnPosition
	destination     gamepoint.GamePointDestination
}

func newScoreControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &TieBreakScoreCountControl{ScoreCountControl: scc}
}

func (t *TieBreakScoreCountControl) SetBallStartTurn(turn turning.TurnPosition) {
	t.ballStartTurn = turn
}

func (t *TieBreakScoreCountControl) SetBallCurrentTurn(turn turning.TurnPosition) {
	t.ballCurrentTurn = turn
}

func (t *TieBreakScoreCountControl) SetDefiantSide(turn turning.TurnPosition) {
	t.defiantSide = turn
}

func (t *TieBreakScoreCountControl) SetDestination(destination gamepoint.GamePointDestination) {
	t.destination = destination
}
