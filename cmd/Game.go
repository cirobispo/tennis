package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

func Game(defiantServing turning.TurnPosition, valueA, valueB int) []gamepoint.GamePointing {
	servingSide := turning.New(turning.TPTurnA)
	ballSide := turning.New(turning.TPTurnA)

	servingSide.AddTurnChangeEvent(func(turn turning.TurnPosition) {
		ballSide.SetBeginningTurn(servingSide.CurrentTurn(), true)
	})

	incA, incB := 0, 0
	points := make([]gamepoint.GamePointing, 0, (valueA+valueB)*2)
	pointsToEquivalence := valueA - valueB
	for i := 0; i < int(math.Abs(float64(pointsToEquivalence))); i++ {
		if valueA > valueB {
			if defiantServing == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
			incA++
		} else {
			if defiantServing == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
			incB++
		}
		servingSide.DoTurn()
	}

	for i := int(math.Abs(float64(pointsToEquivalence))); i < valueA+valueB; i++ {
		if pointsToEquivalence > 0 && servingSide.CurrentTurn() == turning.TPTurnA {
			points = append(points, ServingDoubleFalt()...)
		} else if pointsToEquivalence > 0 && servingSide.CurrentTurn() == turning.TPTurnB {
			points = append(points, gamepoint.NewGamePointAce())
		} else if pointsToEquivalence < 0 && servingSide.CurrentTurn() == turning.TPTurnA {
			points = append(points, gamepoint.NewGamePointAce())
		} else if pointsToEquivalence < 0 && servingSide.CurrentTurn() == turning.TPTurnB {
			points = append(points, ServingDoubleFalt()...)
		}

		servingSide.DoTurn()
	}

	return points
}
