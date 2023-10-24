package main

import (
	"cbtennis/cmd"
	"cbtennis/internal/game"
	"cbtennis/internal/player"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/scoring/tiebreak"
	"cbtennis/internal/turning"
	"fmt"
)

type SDWrapper interface {
	WrapChallenger() (A string, B string)
}

// ServingDescriptorWrapper
type ChallegerSDWrapper struct {
	challenger  player.Challenging
	currentTurn turning.TurnPosition
}

type TextSDWrapper struct {
	valueA, valueB string
	currentTurn    turning.TurnPosition
}

func NewChallengerSDWrapper(challenger player.Challenging, position turning.TurnPosition) ChallegerSDWrapper {
	return ChallegerSDWrapper{
		challenger:  challenger,
		currentTurn: position,
	}
}

func (s ChallegerSDWrapper) WrapChallenger() (A string, B string) {
	A = s.challenger.GetDefiantA().GetName()
	B = s.challenger.GetDefiantB().GetName()

	if s.currentTurn != turning.TPTurnA {
		A = s.challenger.GetDefiantB().GetName()
		B = s.challenger.GetDefiantA().GetName()
	}

	return A, B
}

func NewTextSDWrapper(values [2]string, position turning.TurnPosition) TextSDWrapper {
	return TextSDWrapper{
		valueA:      values[0],
		valueB:      values[1],
		currentTurn: position,
	}
}

func (s TextSDWrapper) WrapChallenger() (A string, B string) {
	A = s.valueA
	B = s.valueB

	if s.currentTurn != turning.TPTurnA {
		A = s.valueB
		B = s.valueA
	}

	return A, B
}

func simulateTieBreak(challenge player.Challenging, g game.GameManager, defiantSide turning.TurnPosition) {
	tiebreak := g.(*game.TieBreak)

	exit := false
	tiebreak.AddFinishedGameEvent(func(servingSide turning.TurnPosition, valueA, valueB int) {
		fmt.Printf("Placar terminado: %d x %d\n", valueA, valueB)
		exit = true
	})

	tiebreak.AddServeTurnChangeEvent(func(turn turning.TurnPosition) {
		sdwServingSide := NewTextSDWrapper([2]string{"deuce", "ad"}, turn)
		lado, _ := sdwServingSide.WrapChallenger()

		fmt.Printf("Servindo no %v\n", lado)
	})

	tiebreak.AddDefiantServingTurnEvent(func(challengerTurn, side turning.TurnPosition) {
		sdwWhoServe := NewChallengerSDWrapper(challenge, challengerTurn)
		sdwServingSide := NewTextSDWrapper([2]string{"deuce", "ad"}, side)

		serveName, _ := sdwWhoServe.WrapChallenger()
		lado, _ := sdwServingSide.WrapChallenger()

		fmt.Printf("%v saca do lado %v\n", serveName, lado)
	})

	tiebreak.AddUpdatePointEvent(func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int) {
		fmt.Printf("%dx%d\n", valueA, valueB)
	})

	tiebreak.StartGame()
	points := cmd.TieBreak(defiantSide, 7, 1)

	for _, p := range points {
		tiebreak.AddPointing(p)
		if exit {
			break
		}
	}

	fmt.Println()
}

func main() {
	scc := tiebreak.NewTieBreakScoreCountControl(7, false)
	challenge := cmd.CreateChallenge()
	defiantSide := turning.TPTurnA
	game := game.NewTieBreak(scc, challenge, defiantSide)
	simulateTieBreak(challenge, game, defiantSide)
}
