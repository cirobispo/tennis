package main

import (
	"cbtennis/internal/game"
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
	"cbtennis/internal/set"
	"cbtennis/internal/turning"
	"fmt"
	"math/rand"
)

func createChallenge() player.Challenging {
	teamA := player.NewTeam([]string{"Ciro", "Leo"})
	teamB := player.NewTeam([]string{"Jailson", "Ibrahim"})

	return player.NewChallenge(teamA, teamB)
}

func gamePointDescription(gpType gamepoint.GamePointType) string {
	switch gpType {
	case gamepoint.GPTAce:
		return "Ace"
	case gamepoint.GPTOut:
		return "Out"
	case gamepoint.GPTNet:
		return "Net"
	case gamepoint.GPTWinner:
		return "Winner"
	default:
		return "???"
	}
}

func simulateGame(game game.GameManager, points []gamepoint.GamePointing) {
	game.AddUpdatePointEvent(func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int) {
		lado := "sacador"
		if turn != turning.TPBegin {
			lado = "recebedor"
		}
		helper := scoring.NewScoreDataWrapper(game.GetScore().GetScoringType(), valueA, valueB)
		vA := helper.GetValueA()
		vB := helper.GetValueB()
		fmt.Printf("%s efetua %v. Placar: %s x %s\n", lado, gamePointDescription(point), vA, vB)
	})

	exit := false
	game.AddFinishedGameEvent(func(valueA, valueB int) {
		helper := scoring.NewScoreDataWrapper(game.GetScore().GetScoringType(), valueA, valueB)

		fmt.Printf("Placar terminado: %s x %s", helper.GetValueA(), helper.GetValueB())
		exit = true
	})

	for _, p := range points {
		game.AddPointing(p)
		if exit {
			break
		}
	}

	fmt.Println()
}

func simulaSet(set *set.Set) {
	sair := false
	set.AddStartingNewGameEvent(func() {
		fmt.Println("Iniciando nova partida")
	})

	set.AddUpdatePointEvent(func(valueA, valueB int) {
		fmt.Printf("Set %d x %d\n\n", valueA, valueB)
	})

	set.AddFinishedSetEvent(func(valueA, valueB int) {
		fmt.Printf("Set encerrado: %d x %d\n", valueA, valueB)
		sair = true
	})

	for {
		game := set.NewGame()

		var points []gamepoint.GamePointing
		if r := rand.Intn(2); r == 0 {
			points = servingWins()
		} else {
			points = receivingWins()
		}

		simulateGame(game, points)
		set.AddGame(game)

		if sair {
			break
		}
	}
}

func main() {
	// scc := gamescore.NewGameScoreCountControl(4, true)
	// game := game.NewSingleStandardGame(scc, createChallenge)

	set := set.New(4, true, false)
	simulaSet(set)
}
