package main

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math/rand"
)

type simpleDevolution int

const (
	pointNet    simpleDevolution = simpleDevolution(gamepoint.GPTNet)
	pointOut    simpleDevolution = simpleDevolution(gamepoint.GPTOut)
	pointIn     simpleDevolution = simpleDevolution(gamepoint.GPTIn)
	pointWinner simpleDevolution = simpleDevolution(gamepoint.GPTWinner)
)

func devolution(dev simpleDevolution, turn turning.TurnPosition) gamepoint.GamePointing {
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

func servingInReturningOut() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	serveIn := gamepoint.NewGamePointServeIn()
	points = append(points, serveIn)

	returnIn := gamepoint.NewGamePointReturnOut()
	points = append(points, returnIn)

	return points
}

func servingInReturningIn() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	serveIn := gamepoint.NewGamePointServeIn()
	points = append(points, serveIn)

	returnIn := gamepoint.NewGamePointReturnIn()
	points = append(points, returnIn)

	return points
}

func servingInReturningNet() []gamepoint.GamePointing {
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
		points = append(points, devolution(pointIn, turn.CurrentTurn()))
		turn.DoTurn()
	}

	if len(points) > 0 {
		points = points[0 : len(points)-1]
		turn.DoTurn()
	}

	beginAligned := len(points)%2 == 0 && turn.CurrentTurn() == turning.TPTurnA
	oppositeAligned := len(points)%2 == 1 && turn.CurrentTurn() == turning.TPTurnB

	if beginAligned || (oppositeAligned && tp == turning.TPTurnB) {
		points = append(points, devolution(pointWinner, tp))
	} else {
		var dev simpleDevolution = pointNet
		if rand.Intn(2) == 1 {
			dev = pointOut
		}
		points = append(points, devolution(dev, turn.CurrentTurn()))
	}

	return points
}

func rallyWinningByServing(hits int) []gamepoint.GamePointing {
	return rallyWinningBy(hits, turning.TPTurnA)
}

func rallyWinningByReceiving(hits int) []gamepoint.GamePointing {
	return rallyWinningBy(hits, turning.TPTurnB)
}

func servingWins() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	points = append(points, gamepoint.NewGamePointAce())

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	return points
}

func receivingWins() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, gamepoint.NewGamePointAce())

	points = append(points, servingInReturningNet()...)
	points = append(points, servingInReturningNet()...)

	points = append(points, servingInReturningOut()...)
	points = append(points, servingInReturningOut()...)

	return points
}

func tiebreakAWins(pointCount int) []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, pointCount)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByReceiving(1)...)

	points = append(points, servingInReturningIn()...)
	points = append(points, rallyWinningByServing(1)...)

	return points
}
