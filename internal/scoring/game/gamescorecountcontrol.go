package game

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
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

func (c GameScoreCountControl) IsDone(valueA, valueB int) bool {
	if c.UpdateHandler != nil {
		return c.IsDoneHandler(c.MaxValue(), c.HasToConfirm(), valueA, valueB)
	}
	return false
}

func (c GameScoreCountControl) UpdateScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	gscc := scc.(*GameScoreCountControl)

	if c.UpdateHandler != nil {
		c.UpdateHandler(gscc, valueA, valueB)
	}
}
