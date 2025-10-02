package turning25

type TurnSide int

const (
	TSA TurnSide = 0
	TSB TurnSide = 1
)

type OnSideChange func(turn TurnSide, undo bool)
type OnTurnRestart func(turn TurnSide)

type Turning interface {
	Restart(executeTurnChangeEvent bool)
	CurrentSide() TurnSide
	StartingSide() TurnSide
	Do()
	Undo()
}

type TurningEvents interface {
	AddResetTurnEvent(resetTurnEvent OnTurnRestart)
	AddTurnChangeEvent(turnChangeEvent OnSideChange)
}
