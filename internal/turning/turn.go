package turning

type TurnPosition int

const (
	TPTurnA TurnPosition = 0
	TPTurnB TurnPosition = 1
)

type OnTurnChange func(turn TurnPosition)
type OnTurnReset func(turn TurnPosition)

type Turn struct {
	currentTurn   TurnPosition
	beginningTurn TurnPosition

	resetTurnEvent  []OnTurnReset
	changeTurnEvent []OnTurnChange
}

func New(turnStartPosition TurnPosition) *Turn {
	return &Turn{
		currentTurn:     turnStartPosition,
		beginningTurn:   turnStartPosition,
		resetTurnEvent:  make([]OnTurnReset, 0),
		changeTurnEvent: make([]OnTurnChange, 0),
	}
}

func (t *Turn) ResetTurn(executeTurnChangeEvent bool) {
	t.currentTurn = t.beginningTurn
	t.executeOnTurnReset()

	if executeTurnChangeEvent {
		t.executeOnTurnChange()
	}
}

func (t Turn) CurrentTurn() TurnPosition {
	return t.currentTurn
}

func (t Turn) BeginningTurn() TurnPosition {
	return t.beginningTurn
}

func (t *Turn) SetBeginningTurn(turnStartPosition TurnPosition, resetTurn bool) {
	t.beginningTurn = turnStartPosition
	if resetTurn {
		t.ResetTurn(false)
	}
}

func (t *Turn) AddResetTurnEvent(resetTurnEvent OnTurnReset) {
	t.resetTurnEvent = append(t.resetTurnEvent, resetTurnEvent)
}

func (t *Turn) AddTurnChangeEvent(turnChangeEvent OnTurnChange) {
	t.changeTurnEvent = append(t.changeTurnEvent, turnChangeEvent)
}

func (t *Turn) DoTurn() {
	if t.currentTurn != -1 {
		if t.currentTurn < 1 {
			t.currentTurn++
		} else {
			t.currentTurn = 0
		}

		if t.changeTurnEvent != nil {
			t.executeOnTurnChange()
		}
	}
}

func (t Turn) executeOnTurnChange() {
	for i := 0; i < len(t.changeTurnEvent); i++ {
		evt := t.changeTurnEvent[i]
		evt(t.currentTurn)
	}
}

func (t Turn) executeOnTurnReset() {
	for i := 0; i < len(t.resetTurnEvent); i++ {
		evt := t.resetTurnEvent[i]
		evt(t.beginningTurn)
	}
}
