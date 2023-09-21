package tennisstatus

import (
	"fmt"
)

type OnGameStart func()
type OnUpdatePoint func(increasedPoint bool)
type OnGameFinish func()

type GameManager interface {
	StartGame()
	AddPointing(point GamePointing)
	AddGameStartEvent(gameStartEvent OnGameStart)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddGameFinishEvent(gameFinishEvent OnGameFinish)
}

type StandardGame struct {
	side   TurnManager
	match  *MatchManager
	score  ScoreManager
	points []GamePointing

	gameStartEvent   []OnGameStart
	updatePointEvent []OnUpdatePoint
	gameFinishEvent  []OnGameFinish
}

func NewStandardGame(match *MatchManager) StandardGame {
	turn := NewTurnManager()
	score := NewGameScore()

	return StandardGame{
		side:             turn,
		match:            match,
		score:            &score,
		points:           make([]GamePointing, 0),
		gameStartEvent:   make([]OnGameStart, 0),
		updatePointEvent: make([]OnUpdatePoint, 0),
		gameFinishEvent:  make([]OnGameFinish, 0),
	}
}

func NewStandardGameInSet(set *StandardSet) StandardGame {
	return NewStandardGame(set.match)
}

func (g *StandardGame) StartGame() {
	onTurnStart := func() {
		fmt.Printf("turno iniciado\n")
	}

	onTurnChange := func(turnIndex int) {
		turnName := "Ad in"
		if turnIndex == 1 {
			turnName = "Ad out"
		}
		fmt.Printf("Serve: %s\t", turnName)
	}
	g.side.AddTurnChangeEvent(onTurnChange)
	g.side.AddTurnStartEvent(onTurnStart)

	scoreGameEvent := func(v1, v2 int) {
		g.executeGameFinishEvent()
	}

	g.score.AddReachedScoreEvent(scoreGameEvent)

	g.side.StartTurn()

	if g.gameStartEvent != nil {
		g.executeGameStartEvent()
	}
}

func (g *StandardGame) AddPointing(point GamePointing) {
	g.points = append(g.points, point)

	ga := NewGameAction(g.score)
	ga.ExecuteAction(point, g.side.turnIndex)

	if g.updatePointEvent != nil {
		g.executeUpdatePointEvent(point.UpdateScore())
	}

	g.side.Do()
}

func (g *StandardGame) AddGameStartEvent(gameStartEvent OnGameStart) {
	g.gameStartEvent = append(g.gameStartEvent, gameStartEvent)
}

func (g *StandardGame) AddUpdatePointEvent(updatePointEvent OnUpdatePoint) {
	g.updatePointEvent = append(g.updatePointEvent, updatePointEvent)
}

func (g *StandardGame) AddGameFinishEvent(gameFinishEvent OnGameFinish) {
	g.gameFinishEvent = append(g.gameFinishEvent, gameFinishEvent)
}

func (g StandardGame) executeGameStartEvent() {
	for i := 0; i < len(g.gameStartEvent); i++ {
		evt := g.gameStartEvent[i]
		evt()
	}
}

func (g StandardGame) executeGameFinishEvent() {
	for i := 0; i < len(g.gameFinishEvent); i++ {
		evt := g.gameFinishEvent[i]
		evt()
	}
}

func (g StandardGame) executeUpdatePointEvent(increasedPoint bool) {
	for i := 0; i < len(g.updatePointEvent); i++ {
		evt := g.updatePointEvent[i]
		evt(increasedPoint)
	}
}
