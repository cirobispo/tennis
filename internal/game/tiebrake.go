package game

import (
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/scoring/tiebreak"
	"cbtennis/internal/turning"
)

type OnDefiantServingTurn func(challengerTurn, side turning.TurnPosition)

type TieBreak struct {
	*StandardGame
	defiantTurn       *turning.Turn
	turnshiftCounting int

	defiantServingTurnEvent []OnDefiantServingTurn
}

func newTieBreak(scc scoring.ScoringCountControl, challenge player.Challenging, startSide turning.TurnPosition) *TieBreak {
	game := &TieBreak{
		StandardGame: newGame(scc, challenge, startSide),
		defiantTurn:  turning.New(startSide),
	}

	game.defiantTurn.AddTurnChangeEvent(func(turn turning.TurnPosition) {
		game.executeDefiantServingTurnChange(turn, game.servingSide.CurrentTurn())
	})

	return game
}

func (t *TieBreak) StartGame() {
	t.StandardGame.StartGame()
	t.defiantTurn.ResetTurn(true)
	t.turnshiftCounting = 1
}

func (t *TieBreak) updateTSCCData(point gamepoint.GamePointing) *tiebreak.TieBreakScoreCountControl {
	tscc := t.score.GetScoreCountControl().(*tiebreak.TieBreakScoreCountControl)
	tscc.SetBallStartTurn(t.ballSide.BeginningTurn())
	tscc.SetBallCurrentTurn(t.ballSide.CurrentTurn())
	tscc.SetDestination(point.PointDestination())
	tscc.SetDefiantSide(t.defiantTurn.CurrentTurn())

	return tscc
}

func (t *TieBreak) AddPointing(point gamepoint.GamePointing) {
	t.points = append(t.points, point)

	pointAdded := point.UpdateScore() == gamepoint.GPUYes || t.isThereDoubleFault()
	if pointAdded {
		tscc := t.updateTSCCData(point)
		t.score.UpdateScore()
		if isDone := tscc.IsDone(t.score.GetStatus()); !isDone {
			t.servingSide.DoTurn()
			t.ballSide.SetBeginningTurn(t.servingSide.CurrentTurn(), true)
			t.turnshiftCounting++
			if t.turnshiftCounting == 2 {
				t.defiantTurn.DoTurn()
				t.turnshiftCounting = 0
			}
		}
	} else {
		if point.UpdateScore() == gamepoint.GPUNo && point.PointDestination() == gamepoint.GPDNone && point.GetType() != gamepoint.GPTServeLet {
			t.ballSide.DoTurn()
		}
	}
}

func (t *TieBreak) AddDefiantServingTurnEvent(challengerTurnEvent OnDefiantServingTurn) {
	t.defiantServingTurnEvent = append(t.defiantServingTurnEvent, challengerTurnEvent)
}

func (t TieBreak) executeDefiantServingTurnChange(challengerTurn, side turning.TurnPosition) {
	for _, evt := range t.defiantServingTurnEvent {
		evt(challengerTurn, side)
	}
}
