package tennisstatus

type OnGameStart func()
type OnUpdatePoint func(increasedPoint bool)
type OnGameFinish func()

type GameManager interface {
	GetScore() ScoreManager
	GetScoreData() ScoreData
	AddPointing(point GamePointing)
	AddGameStartEvent(gameStartEvent OnGameStart)
	AddUpdatePointEvent(updatePointEvent OnUpdatePoint)
	AddGameFinishEvent(gameFinishEvent OnGameFinish)
	AddBallTurnChangeEvent(turnChangeEvent OnTurnChange)
}

type StandardGame struct {
	ballSide  TurnManager
	serveSide TurnManager
	match     *TennisMatch
	score     ScoreManager
	points    []GamePointing

	gameStartEvent   []OnGameStart
	updatePointEvent []OnUpdatePoint
	gameFinishEvent  []OnGameFinish
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
		gameStartEvent:   make([]OnGameStart, 0),
		updatePointEvent: make([]OnUpdatePoint, 0),
		gameFinishEvent:  make([]OnGameFinish, 0),
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

func (g StandardGame) GetScoreData() ScoreData {
	return g.match.score
}

func (g *StandardGame) UpdateBallAndServeTurn(AB TurnPosition, pointAdded bool) {
	if pointAdded {
		g.ballSide.ResetTurn()
		g.serveSide.ResetTurn()
	} else {
		if lastPoint := g.points[len(g.points)-1]; lastPoint.GetType() != GPTServeLet {
			g.ballSide.Turn()
		}
	}
}

func (g *StandardGame) AddPointing(point GamePointing) {
	if len(g.points) == 0 {
		g.executeGameStartEvent()
	}

	g.points = append(g.points, point)

	pointAdded := (point.UpdateScore() == GPUYes)
	if point.UpdateScore() == GPUYes {
		g.score.UpdateScore(g.ballSide.turnIndex)
	} else if point.UpdateScore() == GPUCondicional {
		if len(g.points) > 1 {
			if lastPoint := g.points[len(g.points)-2]; lastPoint.UpdateScore() == GPUCondicional {
				g.score.UpdateScore(g.ballSide.turnIndex)
				pointAdded = true
			}
		}
	}

	if g.updatePointEvent != nil {
		g.executeUpdatePointEvent(pointAdded)
	}

	g.UpdateBallAndServeTurn(g.ballSide.CurrentTurn(), pointAdded)
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

func (g StandardGame) executeUpdatePointEvent(increasedPoint bool) {
	for i := 0; i < len(g.updatePointEvent); i++ {
		evt := g.updatePointEvent[i]
		evt(increasedPoint)
	}
}
