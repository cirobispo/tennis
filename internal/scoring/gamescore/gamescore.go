package gamescore

import (
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"math"
)

type GameScore struct {
	*scoring.Score
}

func New(scc scoring.ScoringCountControl) scoring.Scoring {
	result := &GameScore{
		Score: scoring.NewCustomGameScore(scoring.STGame, scc),
	}

	scc.PlugToScoring(result.Score)

	return result
}

func NewGameScoreCountControl(maxValue int, hasToConfirm bool) scoring.ScoringCountControl {
	return newGameScoreCountControl(maxValue, hasToConfirm, updateGameScore, isGameFinished)
}

func updateGameScore(scc scoring.ScoringCountControl, valueA, valueB *int) {
	gcss := scc.(*GameScoreCountControl)

	var increment = 1
	var value *int

	pointOnA := (gcss.turn == turning.TPTurnA && gcss.destination == gamepoint.GPDSameSide) ||
		(gcss.turn == turning.TPTurnB && gcss.destination == gamepoint.GPDOpositeSide)

	if *valueB > scc.MaxValue() || *valueA > scc.MaxValue() {
		increment = -1
	}

	if pointOnA {
		value = valueA
		if *valueB > scc.MaxValue() {
			value = valueB
		}
	} else {
		value = valueB
		if *valueA > scc.MaxValue() {
			value = valueA
		}
	}
	*value += increment
}

func isGameFinished(maxValue int, hasToConfirm bool, valueA, valueB int) bool {
	simpleWin := (valueA >= maxValue && valueB < maxValue-1) || (valueB >= maxValue && valueA < maxValue-1)
	confirmWin := hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > maxValue || valueB > maxValue)

	return simpleWin || confirmWin
}

// if pointDestination == gamepoint.GPDSameSide {
// 	value = valueA
// 	if turn == turning.TPOpposite {
// 		value = valueB

// 		if *valueA > maxValue {
// 			increment = -1
// 			value = valueA
// 		}
// 	} else {
// 		if *valueB > maxValue {
// 			increment = -1
// 			value = valueB
// 		}
// 	}
// } else {
// 	value = valueB
// 	if turn == turning.TPOpposite {
// 		value = valueA

// 		if *valueB > maxValue {
// 			increment = -1
// 			value = valueB
// 		}
// 	} else {
// 		if *valueA > maxValue {
// 			increment = -1
// 			value = valueA
// 		}
// 	}
// }
