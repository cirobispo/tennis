package main

import (
	"cbtennis/cmd"
	"cbtennis/internal/game"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/gamescore"
	"cbtennis/internal/scoring/gamescore/gamepoint"
	"cbtennis/internal/turning"
	"fmt"
)

func simulateGame(game game.GameManager, points []gamepoint.GamePointing) {
	game.AddUpdatePointEvent(func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int) {
		lado := "sacador"
		if turn != turning.TPTurnA {
			lado = "recebedor"
		}
		helper := scoring.NewScoreDataWrapper(game.GetScore().GetScoringType(), valueA, valueB)
		vA := helper.GetValueA()
		vB := helper.GetValueB()
		fmt.Printf("%s efetua %v. Placar: %s x %s\n", lado, cmd.GamePointDescription(point), vA, vB)
	})

	exit := false
	game.AddFinishedGameEvent(func(servingSide turning.TurnPosition, valueA, valueB int) {
		helper := scoring.NewScoreDataWrapper(game.GetScore().GetScoringType(), valueA, valueB)

		fmt.Printf("Placar terminado: %s x %s", helper.GetValueA(), helper.GetValueB())
		exit = true
	})

	game.StartGame()
	for _, p := range points {
		game.AddPointing(p)
		if exit {
			break
		}
	}

	fmt.Println()
}

func main() {
	scc := gamescore.NewGameScoreCountControl(4, true)
	clgr := cmd.CreateChallenge()
	game := game.NewSingleStandardGame(scc, clgr, turning.TPTurnB)
	simulateGame(game, cmd.Score_Gamex10())
}
