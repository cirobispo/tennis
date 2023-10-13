package set

import (
	"cbtennis/internal/game"
	"cbtennis/internal/scoring"
)

type OnStartingNewGame func()
type OnUpdatePoint func(valueA, valueB int)
type OnStartedSet func()
type OnFinnishedSet func(valueA, valueB int)

type Setting interface {
	StartSet()
	NewGame() game.GameManager
	GetScore() scoring.Scoring
	AddGame(game game.GameManager)
	AddStartingNewGameEvent(startingNewGameEvent OnStartingNewGame)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddStartedSetEvent(startedSetEvent OnStartedSet)
	AddFinishedSetEvent(finnishedSetEvent OnFinnishedSet)
}
