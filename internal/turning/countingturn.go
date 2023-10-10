package turning

type CountingTurn struct {
	*Turn
	counting int
}

func CountingTurnWrapper(t *Turn) Turning {
	return &CountingTurn{Turn: t}
}

func (c *CountingTurn) ResetTurn(executeTurnChangeEvent bool) {
	c.counting = 0
	c.Turn.ResetTurn(executeTurnChangeEvent)
}

func (c *CountingTurn) DoTurn() {
	c.Turn.DoTurn()
	c.counting++
}

func (c CountingTurn) Count() int {
	return c.counting
}
