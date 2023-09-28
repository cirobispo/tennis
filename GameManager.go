package cbtennis

type OnGameStarting func()
type OnUpdatePoint func(increasedPoint bool, increasedBy TurnPosition)
type OnFinishedGame func()

type GameManager interface {
	GetScore() ScoreManager
	AddPointing(point GamePointing)
	AddGameStartingEvent(gameStartEvent OnGameStarting)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddFinishedGameEvent(gameFinishEvent OnFinishedGame)
	AddBallTurnChangeEvent(turnChangeEvent OnTurnChange)
}

type StandardGame struct {
	ballSide  TurnManager
	serveSide TurnManager
	match     *TennisMatch
	score     ScoreManager
	points    []GamePointing

	gameStartEvent   []OnGameStarting
	updatePointEvent []OnUpdatePoint
	gameFinishEvent  []OnFinishedGame
}

func NewSingleStandardGame(match *TennisMatch) StandardGame {
	score := NewGameScore()
	ballSide := NewTurnManager(TPEven)
	serveSide := NewTurnManager(TPEven)

	game := StandardGame{
		ballSide:         ballSide,
		serveSide:        serveSide,
		match:            match,
		score:            &score,
		points:           make([]GamePointing, 0),
		gameStartEvent:   make([]OnGameStarting, 0),
		updatePointEvent: make([]OnUpdatePoint, 0),
		gameFinishEvent:  make([]OnFinishedGame, 0),
	}

	score.AddReachedScoreEvent(func(valueA, valueB int) {
		game.executeGameFinishEvent()
	})

	return game
}

func NewGroupedStandardGame(set *StandardSet) StandardGame {
	return NewSingleStandardGame(set.match)
}

func (g StandardGame) GetScore() ScoreManager {
	return g.score
}

func (g *StandardGame) isThereDoubleFault() bool {
	sum := 0

	size := len(g.points)
	for i := size - 1; i >= 0; i-- {
		if item := g.points[i]; item.UpdateScore() == GPUCondicional {
			sum++
		} else if item.GetType() != GPTServeLet {
			break
		}
	}

	return sum == 2
}

func (g *StandardGame) UpdateBallAndServeTurn(AB TurnPosition, point GamePointing, pointAdded bool) {
	if pointAdded {
		g.ballSide.ResetTurn()
	} else {
		g.ballSide.Turn()
	}
}

func (g *StandardGame) AddPointing(point GamePointing) {
	if len(g.points) == 0 {
		g.executeGameStartEvent()
	}

	g.points = append(g.points, point)

	pointAdded := point.UpdateScore() == GPUYes || g.isThereDoubleFault()
	if pointAdded {
		g.score.UpdateScore(g.ballSide.turnReference, point.PointDestination())

		if g.updatePointEvent != nil {
			g.executeUpdatePointEvent(pointAdded, g.ballSide.BeginningTurn())
		}

		// if s.maxPoints.GetToMaxPoint(s.valueA, s.valueB) {
		// 	s.executeScoreGameEvent(s.valueA, s.valueB)
		// }

		g.ballSide.ResetTurn()
	} else {
		if point.UpdateScore() == GPUNo && point.PointDestination() == GPDNone && point.GetType() != GPTServeLet {
			g.ballSide.Turn()
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

func (g *StandardGame) AddBallTurnChangeEvent(turnChangeEvent OnTurnChange) {
	g.ballSide.AddTurnChangeEvent(turnChangeEvent)
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

func (g StandardGame) executeUpdatePointEvent(increasedPoint bool, increasedBy TurnPosition) {
	for i := 0; i < len(g.updatePointEvent); i++ {
		evt := g.updatePointEvent[i]
		evt(increasedPoint, increasedBy)
	}
}
