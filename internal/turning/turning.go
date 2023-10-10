package turning

type Turning interface {
	ResetTurn(executeTurnChangeEvent bool)
	CurrentTurn() TurnPosition
	BeginningTurn() TurnPosition
	AddResetTurnEvent(resetTurnEvent OnTurnReset)
	AddTurnChangeEvent(turnChangeEvent OnTurnChange)
	DoTurn()
}
