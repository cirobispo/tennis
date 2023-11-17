package cmd

import (
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

func TieBreak(defiantStarting turning.TurnPosition, valueA, valueB int) []gamepoint.GamePointing {
	servingCount := 1
	defiantServing := turning.New(defiantStarting)
	ballSide := turning.New(defiantStarting)

	incA, incB := 0, 0
	pInc := &incB
	whoLoses := turning.TPTurnA
	if valueA > valueB {
		whoLoses = turning.TPTurnB
		pInc = &incA
	}

	servingSide := turning.New(turning.TPTurnA)
	servingSide.AddTurnChangeEvent(func(turn turning.TurnPosition) {
		ballSide.SetBeginningTurn(defiantServing.CurrentTurn(), true)

		servingCount++
		if servingCount == 2 {
			defiantServing.DoTurn()
			servingCount = 0
		}
	})

	defiantServing.AddTurnChangeEvent(func(turn turning.TurnPosition) {
		if whoLoses == turning.TPTurnA {
			whoLoses = turning.TPTurnB
		} else {
			whoLoses = turning.TPTurnA
		}
	})

	points := make([]gamepoint.GamePointing, 0, (valueA+valueB)*2)
	pointsToEquivalence := valueA - valueB

	for i := 0; i < int(math.Abs(float64(pointsToEquivalence))); i++ {
		points = append(points, ServingInPointOutBy(defiantStarting, whoLoses, 2)...)
		*pInc++

		servingSide.DoTurn()
	}
	/**
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

/**
	for i := 0; i < int(math.Abs(float64(pointsToEquivalence))); i++ {
		if valueA > valueB {
			if defiantServing.CurrentTurn() == turning.TPTurnA {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				// points = append(points, ServingDoubleFalt()...)
				whoLoses := turning.TPTurnB
				points = append(points, ServingInPointOutBy(defiantServing.CurrentTurn(), whoLoses, 2)...)
			}
			incA++
		} else {
			if defiantServing.CurrentTurn() == turning.TPTurnB {
				points = append(points, gamepoint.NewGamePointAce())
			} else {
				// points = append(points, ServingDoubleFalt()...)
				whoLoses := turning.TPTurnA
				points = append(points, ServingInPointOutBy(defiantServing.CurrentTurn(), whoLoses, 2)...)
			}
			incB++
		}
		servingSide.DoTurn()
	}
/**/
