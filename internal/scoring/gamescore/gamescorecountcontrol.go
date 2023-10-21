package gamescore

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
)

type GameScoreCountControl struct {
	*scoring.ScoreCountControl
	turn        turning.TurnPosition
	destination gamepoint.GamePointDestination
}

func newGameScoreCountControl(maxValue int, confirm bool, update scoring.UpdatingScoreHandler, done scoring.ScoreIsDoneHandler) scoring.ScoringCountControl {
	scc := scoring.NewScoreCountControl(maxValue, confirm, update, done)
	return &GameScoreCountControl{ScoreCountControl: scc}
}

func (c *GameScoreCountControl) SetTurn(turn turning.TurnPosition) {
	c.turn = turn
}

func (c *GameScoreCountControl) SetDestination(destination gamepoint.GamePointDestination) {
	c.destination = destination
}
