package main

import (
	"cbtennis/internal/game"
	"cbtennis/internal/player"
	"cbtennis/internal/scoring"
	"cbtennis/internal/scoring/game/gamepoint"
	"cbtennis/internal/scoring/tiebreak"
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
	case gamepoint.GPTReturnNet:
		return "net return"
	case gamepoint.GPTReturnOut:
		return "out return"
	default:
		return "???"
	}
}

func simulateGame(game game.GameManager, points []gamepoint.GamePointing) {
	game.AddUpdatePointEvent(func(turn turning.TurnPosition, point gamepoint.GamePointType, valueA, valueB int) {
		lado := "sacador"
		if turn != turning.TPTurnA {
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

	game.StartGame()
	for _, p := range points {
		game.AddPointing(p)
		if exit {
			break
		}
	}

	fmt.Println()
}

func simulateTieBreak(g game.GameManager) {
	tiebreak := g.(*game.TieBreak)

	exit := false
	tiebreak.AddFinishedGameEvent(func(valueA, valueB int) {
		helper := scoring.NewScoreDataWrapper(tiebreak.GetScore().GetScoringType(), valueA, valueB)

		fmt.Printf("Placar terminado: %s x %s", helper.GetValueA(), helper.GetValueB())
		exit = true
	})

	tiebreak.AddServeTurnChangeEvent(func(turn turning.TurnPosition) {
		server := "A"
		if turn == turning.TPTurnB {
			server = "B"
		}

		fmt.Printf("saca do lado %v\n", server)
	})

	tiebreak.AddChallengerTurnChangeEvent(func(challengerTurn, side turning.TurnPosition) {
		server := "A"
		if challengerTurn == turning.TPTurnB {
			server = "B"
		}

		pos := "deuce"
		if side == turning.TPTurnB {
			pos = "ad"
		}

		fmt.Printf("%v saca no lado %v\n", server, pos)
	})

	tiebreak.StartGame()
	points := tiebreakAWins(7)
	for _, p := range points {
		tiebreak.AddPointing(p)
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

	set.StartSet()
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

func simulaMatch() {

}

func main() {
	// scc := gamescore.NewGameScoreCountControl(4, true)
	// game := game.NewSingleStandardGame(scc, createChallenge)

	scc := tiebreak.NewTieBreakScoreCountControl(7, true)
	game := game.NewTieBreak(scc, createChallenge, turning.TPTurnB)
	simulateTieBreak(game)

	// set := set.New(turning.TPTurnA, 4, true, false, true)
	// simulaSet(set)
}
