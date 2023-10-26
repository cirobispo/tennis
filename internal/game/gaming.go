package game

import (
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
)

type OnGameStarting func()
type OnUpdatePoint func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int)
type OnFinishedGame func(servingSide turning.TurnPosition, valueA, valueB int)

type GameType int

const (
	GTGame     GameType = 0
	GTTieBreak GameType = 1
)

type GameManager interface {
	StartGame()
	GetScore() scoring.Scoring
	AddPointing(point gamepoint.GamePointing)
	AddGameStartingEvent(gameStartEvent OnGameStarting)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddFinishedGameEvent(gameFinishEvent OnFinishedGame)
	AddBallTurnChangeEvent(turnChangeEvent turning.OnTurnChange)
}

// Game: scc scoring.ScoringCountControl, challenge player.Challenging, startSide turning.TurnPosition
// TieBreak: scc scoring.ScoringCountControl, challenge player.Challenging, startSide turning.TurnPosition
func New(gameType GameType, scc scoring.ScoringCountControl, challenge player.Challenging, startSide turning.TurnPosition) GameManager {
	if gameType == GTTieBreak {
		return newTieBreak(scc, challenge, startSide)
	}
	return newGame(scc, challenge, startSide)
}
