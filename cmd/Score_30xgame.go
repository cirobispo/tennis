package cmd

import "cbtennis/internal/scoring/gamescore/gamepoint"

func Score_30xGame() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	points = append(points, gamepoint.NewGamePointAce())

	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)

	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)

	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)

	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)

	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)

	return points
}
