package tennisstatus

type TurnInfo interface {
	TurnDescription() string
}

type OnTurnStart func()
type OnTurnChange func(turnIndex int)

type TurnManager struct {
	turnIndex int

	turnStartEvent  []OnTurnStart
	turnChangeEvent []OnTurnChange
}

func NewTurnManager() TurnManager {
	return TurnManager{
		turnIndex:       -1,
		turnStartEvent:  make([]OnTurnStart, 0),
		turnChangeEvent: make([]OnTurnChange, 0),
	}
}

func (t *TurnManager) StartTurn() {
	if t.turnIndex == -1 {
		t.turnIndex = 0
	}

	if t.turnStartEvent != nil {
		t.executeOnTurnStart()
	}
}

func (t TurnManager) executeOnTurnChange(turnIndex int) {
	for i := 0; i < len(t.turnChangeEvent); i++ {
		evt := t.turnChangeEvent[i]
		evt(turnIndex)
	}
}

func (t TurnManager) executeOnTurnStart() {
	for i := 0; i < len(t.turnStartEvent); i++ {
		evt := t.turnStartEvent[i]
		evt()
	}
}

func (t *TurnManager) AddTurnStartEvent(turnStartEvent OnTurnStart) {
	t.turnStartEvent = append(t.turnStartEvent, turnStartEvent)
}

func (t *TurnManager) AddTurnChangeEvent(turnChangeEvent OnTurnChange) {
	t.turnChangeEvent = append(t.turnChangeEvent, turnChangeEvent)
}

func (t *TurnManager) Do() {
	if t.turnIndex != -1 {
		if t.turnChangeEvent != nil {
			t.executeOnTurnChange(t.turnIndex)
		}

		if t.turnIndex < 1 {
			t.turnIndex++
		} else {
			t.turnIndex = 0
		}
	}
}

func (t *TurnManager) Undo() {
	if t.turnIndex != -1 {
		if t.turnIndex > 0 {
			t.turnIndex--
		} else {
			t.turnIndex = 1
		}
	}
}

func (t TurnManager) Current() int {
	result := t.turnIndex
	return result
}
