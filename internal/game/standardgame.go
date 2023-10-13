package game

import (
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"fmt"
)

type StandardGame struct {
	ballSide  *turning.Turn
	serveSide *turning.Turn
	challenge player.Challenging
	score     scoring.Scoring
	points    []gamepoint.GamePointing
	started   bool

	gameStartEvent   []OnGameStarting
	updatePointEvent []OnUpdatePoint
	gameFinishEvent  []OnFinishedGame
}

func NewSingleStandardGame(scc scoring.ScoringCountControl, challenge player.Challenging, startSide turning.TurnPosition) *StandardGame {
	score := gamescore.New(scc)
	ballSide := turning.New(startSide)
	serveSide := turning.New(startSide)

	game := &StandardGame{
		ballSide:         ballSide,
		serveSide:        serveSide,
		challenge:        challenge,
		score:            score,
		points:           make([]gamepoint.GamePointing, 0),
		gameStartEvent:   make([]OnGameStarting, 0),
		updatePointEvent: make([]OnUpdatePoint, 0),
		gameFinishEvent:  make([]OnFinishedGame, 0),
	}

	chScEvent := func(valueA, valueB string) {
		game.executeUpdatePointEvent()
	}
	score.AddChangedScoreEvent(chScEvent)

	score.AddReachedScoreEvent(func(valueA, valueB string) {
		game.executeGameFinishEvent()
	})

	return game
}

func (g *StandardGame) StartGame() {
	fmt.Printf("Jogo iniciado\n")
	g.ballSide.ResetTurn(true)
	g.serveSide.ResetTurn(true)
	g.executeGameStartEvent()
	g.started = true
}

func (g StandardGame) GetScore() scoring.Scoring {
	return g.score
}

func (g *StandardGame) isThereDoubleFault() bool {
	sum := 0

	size := len(g.points)
	for i := size - 1; i >= 0; i-- {
		if item := g.points[i]; item.UpdateScore() == gamepoint.GPUCondicional {
			sum++
		} else if item.GetType() != gamepoint.GPTServeLet {
			break
		}
	}

	return sum == 2
}

func (g *StandardGame) AddPointing(point gamepoint.GamePointing) {
	if g.started {
		g.points = append(g.points, point)

		pointAdded := point.UpdateScore() == gamepoint.GPUYes || g.isThereDoubleFault()
		if pointAdded {
			scc := g.score.GetScoreCountControl().(*gamescore.GameScoreCountControl)
			scc.SetTurn(g.ballSide.CurrentTurn())
			scc.SetDestination(point.PointDestination())
			g.score.UpdateScore()
			g.ballSide.ResetTurn(false)
			g.serveSide.DoTurn()
		} else {
			if point.UpdateScore() == gamepoint.GPUNo && point.PointDestination() == gamepoint.GPDNone && point.GetType() != gamepoint.GPTServeLet {
				g.ballSide.DoTurn()
			}
		}
	}
}

func (g *StandardGame) AddGameStartingEvent(gameStartEvent OnGameStarting) {
	g.gameStartEvent = append(g.gameStartEvent, gameStartEvent)
}

func (g *StandardGame) AddUpdatePointEvent(updatePointEvent OnUpdatePoint) {
	g.updatePointEvent = append(g.updatePointEvent, updatePointEvent)
}

func (g *StandardGame) AddFinishedGameEvent(gameFinishEvent OnFinishedGame) {
	g.gameFinishEvent = append(g.gameFinishEvent, gameFinishEvent)
}

func (g *StandardGame) AddBallTurnChangeEvent(turnChangeEvent turning.OnTurnChange) {
	g.ballSide.AddTurnChangeEvent(turnChangeEvent)
}

func (g *StandardGame) AddServeTurnChangeEvent(turnChangeEvent turning.OnTurnChange) {
	g.serveSide.AddTurnChangeEvent(turnChangeEvent)
}

func (g StandardGame) executeGameStartEvent() {
	for i := 0; i < len(g.gameStartEvent); i++ {
		evt := g.gameStartEvent[i]
		evt()
	}
}

func (g *StandardGame) executeGameFinishEvent() {
	for i := 0; i < len(g.gameFinishEvent); i++ {
		evt := g.gameFinishEvent[i]
		valueA, valueB := g.score.GetStatus()
		evt(valueA, valueB)
	}

	g.started = false
}

func (g StandardGame) executeUpdatePointEvent() {
	for i := 0; i < len(g.updatePointEvent); i++ {
		evt := g.updatePointEvent[i]
		valueA, valueB := g.score.GetStatus()
		lastPoint := g.points[len(g.points)-1]
		evt(g.ballSide.CurrentTurn(), lastPoint.GetType(), valueA, valueB)
	}
}
