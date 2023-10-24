package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

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
			if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, ServingDoubleFalt()...)
			} else if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, ServingDoubleFalt()...)
			}
		} else {
			if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, ServingDoubleFalt()...)
			} else if defiantServing.CurrentTurn() == turning.TPTurnB && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnB {
				points = append(points, ServingDoubleFalt()...)
			} else if defiantServing.CurrentTurn() == turning.TPTurnA && servingSide.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			}
		}

		servingSide.DoTurn()
	}

	return points
}
