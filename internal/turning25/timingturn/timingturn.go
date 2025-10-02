package timingturn

import (
	"cbtennis/internal/turning25/turn"
	"time"
)

type TimingTurn struct {
	*turn.Turn
	startTime time.Time 
}

type Timing interface {
	SlapsedTime() time.Duration	
}

func New(t *turn.Turn) *TimingTurn {
	return &TimingTurn{ Turn: t, startTime: time.Now() }
}

func (t TimingTurn) SlapsedTime() time.Duration {
	return time.Since(t.startTime)
}
