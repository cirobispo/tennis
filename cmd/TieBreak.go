package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

func TieBreak(defiantStarting turning.TurnPosition, valueA, valueB int, hasToConfirm bool) []gamepoint.GamePointing {
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

	incA, incB := 0, 0
	points := make([]gamepoint.GamePointing, 0, max*2)
	pointsToEquivalence := valueA - valueB
	for i := 0; i < int(math.Abs(float64(pointsToEquivalence))); i++ {
		if valueA > valueB {
			if defiantServing.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
			incA++
		} else {
			if defiantServing.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				points = append(points, ServingDoubleFalt()...)
			}
			incB++
		}
		servingSide.DoTurn()
	}
	/**/
	for i := int(math.Abs(float64(pointsToEquivalence))); i < valueA+valueB; i++ {
		if pointsToEquivalence > 0 {
			if (incA - int(math.Abs(float64(pointsToEquivalence)))) >= incB {
				if defiantServing.CurrentTurn() == turning.TPTurnA {
					points = append(points, ServingDoubleFalt()...)
				} else {
					points = append(points, gamepoint.NewGamePointAce())
				}
				incB++
			} else {
				if defiantServing.CurrentTurn() == turning.TPTurnB {
					points = append(points, ServingDoubleFalt()...)
				} else {
					points = append(points, gamepoint.NewGamePointAce())
				}
				incA++
			}
		} else {
			if (incB - int(math.Abs(float64(pointsToEquivalence)))) >= incA {
				if defiantServing.CurrentTurn() == turning.TPTurnB {
					points = append(points, ServingDoubleFalt()...)
				} else {
					points = append(points, gamepoint.NewGamePointAce())
				}
				incA++
			} else {
				if defiantServing.CurrentTurn() == turning.TPTurnA {
					points = append(points, ServingDoubleFalt()...)
				} else {
					points = append(points, gamepoint.NewGamePointAce())
				}
				incB++
			}
		}
		servingSide.DoTurn()
	}
	/**/
	return points
}
