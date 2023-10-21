package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math/rand"
)

type SimpleDevolution int

const (
	pointNet    SimpleDevolution = SimpleDevolution(gamepoint.GPTNet)
	pointOut    SimpleDevolution = SimpleDevolution(gamepoint.GPTOut)
	pointIn     SimpleDevolution = SimpleDevolution(gamepoint.GPTIn)
	pointWinner SimpleDevolution = SimpleDevolution(gamepoint.GPTWinner)
)

func GamePointDescription(gpType gamepoint.GamePointType) string {
	switch gpType {
	case gamepoint.GPTAce:
		return "Ace"
	case gamepoint.GPTOut:
		return "Out"
	case gamepoint.GPTNet:
		return "Net"
	case gamepoint.GPTWinner:
		return "Winner"
	case gamepoint.GPTReturnNet:
		return "net return"
	case gamepoint.GPTReturnOut:
		return "out return"
	default:
		return "???"
	}
}

func Devolution(dev SimpleDevolution, turn turning.TurnPosition) gamepoint.GamePointing {
	switch dev {
	case pointIn:
		return gamepoint.NewGamePointIn()
	case pointNet:
		return gamepoint.NewGamePointNet(turn)
	case pointWinner:
		return gamepoint.NewGamePointWinner()
	default:
		return gamepoint.NewGamePointOut(turn)
	}
}

func ServingDoubleFalt() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	for i := 0; i < 2; i++ {
		if rand.Intn(2) == 0 {
			points = append(points, gamepoint.NewGamePointServeNet())
		} else {
			points = append(points, gamepoint.NewGamePointServeOut())
		}
	}

	return points
}

func ServingInReturningOut() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	serveIn := gamepoint.NewGamePointServeIn()
	points = append(points, serveIn)

	returnIn := gamepoint.NewGamePointReturnOut()
	points = append(points, returnIn)

	return points
}

func ServingInReturningIn() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	serveIn := gamepoint.NewGamePointServeIn()
	points = append(points, serveIn)

	returnIn := gamepoint.NewGamePointReturnIn()
	points = append(points, returnIn)

	return points
}

func ServingInReturningNet() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	serveIn := gamepoint.NewGamePointServeIn()
	points = append(points, serveIn)

	returnIn := gamepoint.NewGamePointReturnNet()
	points = append(points, returnIn)

	return points
}

func rallyWinningBy(hits int, tp turning.TurnPosition) []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	turn := turning.New(tp)

	for i := 0; i < hits; i++ {
		points = append(points, Devolution(pointIn, turn.CurrentTurn()))
		turn.DoTurn()
	}

	if len(points) > 0 {
		points = points[0 : len(points)-1]
		turn.DoTurn()
	}

	beginAligned := len(points)%2 == 0 && turn.CurrentTurn() == turning.TPTurnA
	oppositeAligned := len(points)%2 == 1 && turn.CurrentTurn() == turning.TPTurnB

	if beginAligned || (oppositeAligned && tp == turning.TPTurnB) {
		points = append(points, Devolution(pointWinner, tp))
	} else {
		var dev SimpleDevolution = pointNet
		if rand.Intn(2) == 1 {
			dev = pointOut
		}
		points = append(points, Devolution(dev, turn.CurrentTurn()))
	}

	return points
}

func RallyWinningByServing(hits int) []gamepoint.GamePointing {
	return rallyWinningBy(hits, turning.TPTurnA)
}

func RallyWinningByReceiving(hits int) []gamepoint.GamePointing {
	return rallyWinningBy(hits, turning.TPTurnB)
}
