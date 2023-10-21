package game

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
)

type OnGameStarting func()
type OnUpdatePoint func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int)
type OnFinishedGame func(servingSide turning.TurnPosition, valueA, valueB int)

type GameManager interface {
	StartGame()
	GetScore() scoring.Scoring
	AddPointing(point gamepoint.GamePointing)
	AddGameStartingEvent(gameStartEvent OnGameStarting)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddFinishedGameEvent(gameFinishEvent OnFinishedGame)
	AddBallTurnChangeEvent(turnChangeEvent turning.OnTurnChange)
}
