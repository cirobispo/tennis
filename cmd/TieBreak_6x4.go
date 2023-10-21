package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

func TieBreak_7x4() []gamepoint.GamePointing {
	points := make([]gamepoint.GamePointing, 0, 4)

	points = append(points, gamepoint.NewGamePointAce())
	//*1x0 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)
	//1x*1
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//2x*1 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)
	//*3x1
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)
	//*4x1 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)
	//4x*2
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//5x*2 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//*5x3
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//*5x4 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//6x*4
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByReceiving(1)...)
	//7x*4 v
	points = append(points, ServingInReturningIn()...)
	points = append(points, RallyWinningByServing(1)...)

	return points
}

func TieBreak(defiantStarting turning.TurnPosition, valueA, valueB int) []gamepoint.GamePointing {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}(valueA, valueB)

	servingCount := 1
	defiantServing := turning.New(defiantStarting)

	servingSide := turning.New(turning.TPTurnA)
	servingSide.AddTurnChangeEvent(func(turn turning.TurnPosition) {
		servingCount++
		if servingCount == 2 {
			defiantServing.DoTurn()
			servingCount = 0
		}
	})

	// incA, incB := 0, 0
	points := make([]gamepoint.GamePointing, 0, max*2)
	pointsToEquivalence := valueA - valueB
	for i := 0; i < int(math.Abs(float64(pointsToEquivalence))); i++ {
		if valueA > valueB {
			if defiantServing.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
		} else {
			if defiantServing.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
		}
		servingSide.DoTurn()
	}

	for i := 0; i < valueA+valueB-int(math.Abs(float64(pointsToEquivalence))); i++ {
		if pointsToEquivalence > 0 {
			if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, ServingDoubleFalt()...)
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, ServingDoubleFalt()...)
			}
		} else {
			if defiantServing.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			}
		}

		servingSide.DoTurn()
	}

	return points
}
