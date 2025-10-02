package turn

import (
	turning "cbtennis/internal/turning25"
)

type OnAfterDo func (lastSide turning.TurnSide)

type Turn struct {
	currentSide turning.TurnSide
	startSide   turning.TurnSide

	restartTurnEvent []turning.OnTurnRestart
	changeSideEvent  []turning.OnSideChange
	afterDoEvent     []OnAfterDo
}

func New(turnStartPosition turning.TurnSide) *Turn {
	return &Turn{
		currentSide:      turnStartPosition,
		startSide:        turnStartPosition,
		restartTurnEvent: make([]turning.OnTurnRestart, 0),
		changeSideEvent:  make([]turning.OnSideChange, 0),
		afterDoEvent:     make([]OnAfterDo, 0),
	}
}

func (t *Turn) Restart(executeTurnChangeEvent bool) {
	if executeTurnChangeEvent {
		t.executeOnTurnChange(false)
	}

	t.currentSide = t.startSide
	t.executeOnTurnReset()
}

func (t Turn) CurrentSide() turning.TurnSide {
	return t.currentSide
}

func (t Turn) StartingSide() turning.TurnSide {
	return t.startSide
}

func (t *Turn) SetBeginningTurn(turnStartPosition turning.TurnSide, resetTurn bool) {
	t.startSide = turnStartPosition
	if resetTurn {
		t.Restart(false)
	}
}

func (t *Turn) AddResetTurnEvent(resetTurnEvent turning.OnTurnRestart) {
	t.restartTurnEvent = append(t.restartTurnEvent, resetTurnEvent)
}

func (t *Turn) AddTurnChangeEvent(turnChangeEvent turning.OnSideChange) {
	t.changeSideEvent = append(t.changeSideEvent, turnChangeEvent)
}

func (t *Turn) AddAfterDoEvent(beforeDoEvent OnAfterDo) {
	t.afterDoEvent = append(t.afterDoEvent, beforeDoEvent)
}

func (t *Turn) Do() {
	if t.currentSide < turning.TSB {
		t.currentSide++
	} else {
		t.currentSide = turning.TSA
	}

	t.executeOnAfterDo()
	t.executeOnTurnChange(false)
}

func (t *Turn) Undo() {
	if t.currentSide == turning.TSA {
		t.currentSide = turning.TSB + 1
	}
	
	t.currentSide--

	t.executeOnTurnChange(true)
}

func (t Turn) executeOnTurnChange(undo bool) {
	for i := 0; i < len(t.changeSideEvent); i++ {
		evt := t.changeSideEvent[i]
		evt(t.currentSide, undo)
	}
}

func (t Turn) executeOnTurnReset() {
	for i := 0; i < len(t.restartTurnEvent); i++ {
		evt := t.restartTurnEvent[i]
		evt(t.startSide)
	}
}

func (t Turn) executeOnAfterDo() {
	for i := 0; i < len(t.afterDoEvent); i++ {
		evt := t.afterDoEvent[i]
		evt(t.currentSide)
	}
}
