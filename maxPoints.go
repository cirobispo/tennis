package cbtennis

import "math"

type maxPoints struct {
	maxValue     int
	hasToConfirm bool

	updateScore func(maxValue int, AB TurnPosition, pointDestination GamePointDestination, valueA, valueB *int)
}

func newMaxPoints(max int, confirm bool, updateScore func(maxValue int, AB TurnPosition, pointDestination GamePointDestination, valueA, valueB *int)) maxPoints {
	return maxPoints{
		maxValue:     max,
		hasToConfirm: confirm,
		updateScore:  updateScore,
	}
}

func (p maxPoints) UpdateScore(AB TurnPosition, pointDestination GamePointDestination, valueA, valueB *int) {
	if p.updateScore != nil {
		p.updateScore(p.maxValue, AB, pointDestination, valueA, valueB)
	}
}

func (p maxPoints) NeedsConfirmation() bool {
	return p.hasToConfirm
}

func (p maxPoints) GetToMaxPoint(valueA, valueB int) bool {
	simpleWin := (valueA >= p.maxValue && valueB < p.maxValue-1) || (valueB >= p.maxValue && valueA < p.maxValue-1)
	confirmWin := p.hasToConfirm && int(math.Abs(float64(valueA)-float64(valueB))) == 2 &&
		(valueA > p.maxValue || valueB > p.maxValue)

	return simpleWin || confirmWin
}
