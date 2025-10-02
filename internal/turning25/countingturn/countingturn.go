package countingturn

import (
	turning "cbtennis/internal/turning25"
	"cbtennis/internal/turning25/turn"
)

type CountingTurn struct {
	*turn.Turn
	counting int
}

type Counting interface {
	Count() int
}

func (c *CountingTurn) inc() {
	c.counting++
}

func New(t *turn.Turn) *CountingTurn {
	result := &CountingTurn{Turn: t, counting: 0}
	result.AddAfterDoEvent(func(lastSide turning.TurnSide) {
		result.inc()
	})

	return result
}

func (t *CountingTurn) Restart(executeTurnChangeEvent bool) {
	t.counting = 0
	t.Turn.Restart(executeTurnChangeEvent)
}

func (t CountingTurn) Count() int {
	return t.counting
}
