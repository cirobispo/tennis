package game

import (
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
	"cbtennis/internal/scoring/tiebreak"
	"cbtennis/internal/turning"
)

type OnChallengerServingTurn func(challengerTurn, side turning.TurnPosition)

type TieBreak struct {
	*StandardGame
	challengerSide    *turning.Turn
	turnshiftCounting int

	challengerTurnEvent []OnChallengerServingTurn
}

func NewTieBreak(scc scoring.ScoringCountControl, challenge player.Challenging, challengerSide turning.TurnPosition) *TieBreak {
	game := NewSingleStandardGame(scc, challenge)
	return &TieBreak{
		StandardGame:   game,
		challengerSide: turning.New(challengerSide),
	}
}

func (t *TieBreak) StartGame() {
	t.StandardGame.StartGame()
	t.turnshiftCounting = 1
	t.executeGameStartEvent()
}

func (t *TieBreak) AddPointing(point gamepoint.GamePointing) {
	t.points = append(t.points, point)

	pointAdded := point.UpdateScore() == gamepoint.GPUYes || t.isThereDoubleFault()
	if pointAdded {
		tscc := t.score.GetScoreCountControl().(*tiebreak.TieBreakScoreCountControl)
		tscc.SetTurn(t.ballSide.CurrentTurn())
		tscc.SetDestination(point.PointDestination())
		t.score.UpdateScore()
		t.serveSide.DoTurn()
		t.turnshiftCounting++
		if t.turnshiftCounting == 2 {
			t.challengerSide.DoTurn()
			t.ballSide.SetBeginningTurn(t.challengerSide.CurrentTurn())
			t.executeChallengerTurnChange(t.challengerSide.CurrentTurn(), t.serveSide.CurrentTurn())
			t.turnshiftCounting = 0
		} else {
			t.ballSide.ResetTurn(false)
		}
	} else {
		if point.UpdateScore() == gamepoint.GPUNo && point.PointDestination() == gamepoint.GPDNone && point.GetType() != gamepoint.GPTServeLet {
			t.ballSide.DoTurn()
		}
	}
}

func (t *TieBreak) AddChallengerTurnChangeEvent(challengerTurnEvent OnChallengerServingTurn) {
	t.challengerTurnEvent = append(t.challengerTurnEvent, challengerTurnEvent)
}

func (t TieBreak) executeChallengerTurnChange(challengerTurn, side turning.TurnPosition) {
	for _, evt := range t.challengerTurnEvent {
		evt(challengerTurn, side)
	}
}
