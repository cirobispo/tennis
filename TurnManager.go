package cbtennis

type TurnPosition int

const (
	TPEven TurnPosition = 0
	TPOdd  TurnPosition = 1
)

type OnTurnChange func(turn TurnPosition)
type OnTurnReset func()

type TurnManager struct {
	turnReference     TurnPosition
	turnStartPosition TurnPosition

	resetTurnEvent  []OnTurnReset
	changeTurnEvent []OnTurnChange
}

func NewTurnManager(turnStartPosition TurnPosition) TurnManager {
	return TurnManager{
		turnReference:     turnStartPosition,
		turnStartPosition: turnStartPosition,
		resetTurnEvent:    make([]OnTurnReset, 0),
		changeTurnEvent:   make([]OnTurnChange, 0),
	}
}

type TurnDescriber interface {
	EvenDesc() string
	OddDesc() string
}

type TurnSideDescribe struct {
	even, odd string
}

func NewTurnDescribe(t TurnManager, Even, Odd string) TurnDescriber {
	even := Even
	odd := Odd
	if t.BeginningTurn() != TPEven {
		even = Odd
		odd = Even
	}

	return TurnSideDescribe{even: even, odd: odd}
}

func (d TurnSideDescribe) EvenDesc() string {
	return d.even
}

func (d TurnSideDescribe) OddDesc() string {
	return d.odd
}

func (t *TurnManager) ResetTurn() {
	t.turnReference = t.turnStartPosition
	t.executeOnTurnReset()
}

func (t TurnManager) CurrentTurn() TurnPosition {
	return t.turnReference
}

func (t TurnManager) BeginningTurn() TurnPosition {
	return t.turnStartPosition
}

func (t *TurnManager) AddResetTurnEvent(resetTurnEvent OnTurnReset) {
	t.resetTurnEvent = append(t.resetTurnEvent, resetTurnEvent)
}

func (t *TurnManager) AddTurnChangeEvent(turnChangeEvent OnTurnChange) {
	t.changeTurnEvent = append(t.changeTurnEvent, turnChangeEvent)
}

func (t *TurnManager) Turn() {
	if t.turnReference != -1 {
		if t.changeTurnEvent != nil {
			t.executeOnTurnChange(t.turnReference)
		}

		if t.turnReference < 1 {
			t.turnReference++
		} else {
			t.turnReference = 0
		}
	}
}

func (t TurnManager) executeOnTurnChange(turn TurnPosition) {
	for i := 0; i < len(t.changeTurnEvent); i++ {
		evt := t.changeTurnEvent[i]
		evt(turn)
	}
}

func (t TurnManager) executeOnTurnReset() {
	for i := 0; i < len(t.resetTurnEvent); i++ {
		evt := t.resetTurnEvent[i]
		evt()
	}
}
