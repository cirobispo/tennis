package game

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
	"cbtennis/internal/turning"
)

type OnGameStarting func()
type OnUpdatePoint func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int)
type OnFinishedGame func(valueA, valueB int)

type GameManager interface {
	GetScore() scoring.Scoring
	AddPointing(point gamepoint.GamePointing)
	AddGameStartingEvent(gameStartEvent OnGameStarting)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddFinishedGameEvent(gameFinishEvent OnFinishedGame)
	AddBallTurnChangeEvent(turnChangeEvent turning.OnTurnChange)
}